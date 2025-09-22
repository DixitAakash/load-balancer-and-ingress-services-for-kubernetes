# CA Certificate Cleanup Fix for OpenShift Routes

## Problem Description

When multiple OpenShift Routes share the same CA certificate (due to using the same intermediate CA), AKO encounters a cleanup issue:

1. **Route1** and **Route2** both use the same CA certificate
2. **Avi Controller deduplicates** the CA certificates but AKO doesn't know about this
3. When **Route1** is deleted, AKO tries to delete the CA certificate
4. **Deletion fails** with 403 error: "Cannot delete, object is referred by: ['SSLKeyAndCertificate route2', 'VirtualService route2']"
5. **CA certificate is never cleaned up** even when Route2 is also deleted later

## Root Cause

- AKO uses the **intended CA certificate name** from its own model for cache calculations
- But the **Avi Controller may deduplicate CA certificates** and use different names
- AKO's cache becomes **out of sync** with the actual CA certificate references on Avi Controller
- AKO doesn't have proper **reference counting** for CA certificates based on actual usage

## Solution

### 1. Enhanced SSL Cache Structure

**File**: `internal/cache/cache_utils.go`

Added `ActualCACertName` field to track the actual CA certificate name used by Avi Controller:

```go
type AviSSLCache struct {
    Name               string
    Tenant             string
    Uuid               string
    CloudConfigCksum   uint32
    LastModified       string
    InvalidData        bool
    Cert               string
    HasCARef           bool
    CACertUUID         string
    IntendedCACertName string // NEW: The CA certificate name AKO originally intended to use
    ActualCACertName   string // NEW: The actual CA certificate name used by Avi Controller
    HasReference       bool
}
```

### 2. Dual CA Certificate Name Tracking

**File**: `internal/rest/ssl_key_certificate.go`

Modified `AviSSLKeyCertAdd` to track both intended and actual CA certificate names:

- **Intended CA Certificate Name**: From AKO's original request (`originalSSLObj.CaCerts[0].Name`)
- **Actual CA Certificate Name**: From Avi Controller's response (`resp["ca_certs"][0]["ca_ref"]`)

```go
// Extract intended CA certificate name from AKO's original request
if len(originalSSLObj.CaCerts) > 0 {
    if originalSSLObj.CaCerts[0].Name != nil {
        intendedCACert = *originalSSLObj.CaCerts[0].Name
        hasCA = true
    }
}

// Extract actual CA certificate name from Avi Controller response
if hasCA {
    if caCertsInterface, exists := resp["ca_certs"]; exists {
        // Parse actual CA certificate name from response
        actualCACert = extractCACertNameFromResponse(caCertsInterface)
    }
}

ssl_cache_obj := avicache.AviSSLCache{
    // Use intended CA certificate name for checksum calculation (for cache consistency)
    CloudConfigCksum:   lib.SSLKeyCertChecksum(name, cert, intendedCACert, ...),
    IntendedCACertName: intendedCACert, // Store intended CA certificate name
    ActualCACertName:   actualCACert,   // Store actual CA certificate name for proper cleanup
}
```

### 3. Reference-Counting Based CA Certificate Deletion

**File**: `internal/rest/dequeue_nodes.go`

Enhanced `SSLKeyCertDelete` with proper CA certificate reference counting:

```go
func (rest *RestOperations) SSLKeyCertDelete(ssl_to_delete []avicache.NamespaceName, namespace string, rest_ops []*utils.RestOp, key string) []*utils.RestOp {
    // Build a reference count map of actual CA certificates still in use
    actualCACertRefCount := make(map[string]int)
    
    // Count references from all SSL certificates that are NOT being deleted
    allSSLKeys := rest.cache.SSLKeyCache.AviGetAllKeys()
    for _, sslKey := range allSSLKeys {
        if sslCacheObj.HasCARef && sslCacheObj.ActualCACertName != "" {
            actualCACertRefCount[sslCacheObj.ActualCACertName]++
        }
    }
    
    // Only delete CA certificates that have zero references
    for _, del_ssl := range ssl_to_delete {
        if ssl_cache_obj.HasCARef && ssl_cache_obj.ActualCACertName != "" {
            if actualCACertRefCount[ssl_cache_obj.ActualCACertName] == 0 {
                // Schedule CA certificate for deletion
                caCertsToDelete = append(caCertsToDelete, ssl_cache_obj.ActualCACertName)
            }
        }
    }
}
```

## How the Fix Works

### Before Fix (Broken Scenario)

1. **Route1** created → AKO posts SSL cert with CA ref: `route1-cacert`
2. **Route2** created → AKO posts SSL cert with CA ref: `route2-cacert`
3. **Avi Controller deduplicates** → Both SSL certs actually use `route1-cacert`
4. **AKO's cache** still thinks Route2 uses `route2-cacert`
5. **Route1 deleted** → AKO tries to delete `route1-cacert` → **FAILS** (still referenced by Route2)
6. **Route2 deleted** → AKO tries to delete `route2-cacert` → **SUCCEEDS** (but wrong CA cert)
7. **Result**: `route1-cacert` is never cleaned up

### After Fix (Working Scenario)

1. **Route1** created → AKO posts SSL cert with CA ref: `route1-cacert`
2. **Route2** created → AKO posts SSL cert with CA ref: `route2-cacert`
3. **Avi Controller deduplicates** → Both SSL certs actually use `route1-cacert`
4. **AKO's cache** stores:
   - Route1: `ActualCACertName = "route1-cacert"`
   - Route2: `ActualCACertName = "route1-cacert"` (extracted from response)
5. **Route1 deleted** → AKO counts refs to `route1-cacert` → **Still 1 reference** (Route2) → **Skip deletion**
6. **Route2 deleted** → AKO counts refs to `route1-cacert` → **0 references** → **Delete successfully**
7. **Result**: `route1-cacert` is properly cleaned up

## Benefits

1. **Fixes cache corruption** by using intended CA cert name for checksums
2. **Prevents failed deletions** by using actual CA cert names for cleanup
3. **Proper reference counting** ensures CA certificates are only deleted when safe
4. **Maintains backward compatibility** with existing functionality
5. **Handles Avi Controller deduplication** transparently

## Testing

The fix should be tested with:
1. Multiple OpenShift Routes using the same intermediate CA certificate
2. Deleting routes in different orders
3. Verifying CA certificates are properly cleaned up
4. Ensuring no 403 "object is referred by" errors occur

## Files Modified

- `internal/cache/cache_utils.go` - Added `ActualCACertName` field
- `internal/rest/ssl_key_certificate.go` - Dual CA certificate name tracking
- `internal/rest/dequeue_nodes.go` - Reference-counting based deletion logic
