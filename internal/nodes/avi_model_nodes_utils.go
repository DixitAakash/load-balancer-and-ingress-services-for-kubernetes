/* 
* Copyright 2022-2023 VMware, Inc.
* All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

// Code generated by a tool; DO NOT EDIT.

package nodes

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"

	// avimodels "github.com/vmware/alb-sdk/go/models"

	"github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/apis/ako/v1alpha2"
	"github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/utils"
)

type AviVsNodeGeneratedFields struct {
	Fqdn          *string
	OauthVsConfig *v1alpha2.OAuthVSConfig
	SamlSpConfig  *v1alpha2.SAMLSPConfig
	SsoPolicyRef  *string
}

func (v *AviVsNodeGeneratedFields) CalculateCheckSumOfGeneratedCode() uint32 {
	checksumStringSlice := make([]string, 0, 4)
	if v.Fqdn != nil {
		checksumStringSlice = append(checksumStringSlice, *v.Fqdn)
	}

	if v.OauthVsConfig != nil {
		checksumStringSlice = append(checksumStringSlice, utils.Stringify(v.OauthVsConfig))
	}

	if v.SamlSpConfig != nil {
		checksumStringSlice = append(checksumStringSlice, utils.Stringify(v.SamlSpConfig))
	}

	if v.SsoPolicyRef != nil {
		checksumStringSlice = append(checksumStringSlice, *v.SsoPolicyRef)
	}

	chksumStr := strings.Join(checksumStringSlice, delim)
	checksum := utils.Hash(chksumStr)
	return checksum
}

func (o *AviVsNodeGeneratedFields) ConvertToRef() {
	if o != nil {
		if o.OauthVsConfig != nil {
			for ii := range o.OauthVsConfig.OauthSettings {
				if o.OauthVsConfig.OauthSettings[ii] != nil {
					if o.OauthVsConfig.OauthSettings[ii].AuthProfileRef != nil {
						o.OauthVsConfig.OauthSettings[ii].AuthProfileRef = proto.String("/api/authprofile?name=" + *o.OauthVsConfig.OauthSettings[ii].AuthProfileRef)
					}
				}
			}
		}
		if o.SamlSpConfig != nil {
			if o.SamlSpConfig.SigningSslKeyAndCertificateRef != nil {
				o.SamlSpConfig.SigningSslKeyAndCertificateRef = proto.String("/api/sslkeyandcertificate?name=" + *o.SamlSpConfig.SigningSslKeyAndCertificateRef)
			}
		}
		if o.SsoPolicyRef != nil {
			o.SsoPolicyRef = proto.String("/api/ssopolicy?name=" + *o.SsoPolicyRef)
		}
	}
}

func (o *AviVsNodeCommonFields) ConvertToRef() {
	if o != nil {
		if o.AnalyticsProfileRef != nil {
			o.AnalyticsProfileRef = proto.String("/api/analyticsprofile?name=" + *o.AnalyticsProfileRef)
		}
		if o.ApplicationProfileRef != nil {
			o.ApplicationProfileRef = proto.String("/api/applicationprofile?name=" + *o.ApplicationProfileRef)
		}
		if o.SslProfileRef != nil {
			o.SslProfileRef = proto.String("/api/sslprofile?name=" + *o.SslProfileRef)
		}
		VsDatascriptRefs := make([]string, 0, len(o.VsDatascriptRefs))
		for i := range o.VsDatascriptRefs {
			ref := fmt.Sprintf("/api/vsdatascriptset?name=" + o.VsDatascriptRefs[i])
			if !utils.HasElem(VsDatascriptRefs, ref) {
				VsDatascriptRefs = append(VsDatascriptRefs, ref)
			}
		}
		o.VsDatascriptRefs = VsDatascriptRefs
		if o.WafPolicyRef != nil {
			o.WafPolicyRef = proto.String("/api/wafpolicy?name=" + *o.WafPolicyRef)
		}
	}
}

type AviPoolGeneratedFields struct {
}

func (v *AviPoolGeneratedFields) CalculateCheckSumOfGeneratedCode() uint32 {
	checksumStringSlice := make([]string, 0, 0)
	chksumStr := strings.Join(checksumStringSlice, delim)
	checksum := utils.Hash(chksumStr)
	return checksum
}

func (o *AviPoolGeneratedFields) ConvertToRef() {
}

func (o *AviPoolCommonFields) ConvertToRef() {
	if o != nil {
		if o.ApplicationPersistenceProfileRef != nil {
			o.ApplicationPersistenceProfileRef = proto.String("/api/applicationpersistenceprofile?name=" + *o.ApplicationPersistenceProfileRef)
		}
		HealthMonitorRefs := make([]string, 0, len(o.HealthMonitorRefs))
		for i := range o.HealthMonitorRefs {
			ref := fmt.Sprintf("/api/healthmonitor?name=" + o.HealthMonitorRefs[i])
			if !utils.HasElem(HealthMonitorRefs, ref) {
				HealthMonitorRefs = append(HealthMonitorRefs, ref)
			}
		}
		o.HealthMonitorRefs = HealthMonitorRefs
		if o.PkiProfileRef != nil {
			o.PkiProfileRef = proto.String("/api/pkiprofile?name=" + *o.PkiProfileRef)
		}
		if o.SslKeyAndCertificateRef != nil {
			o.SslKeyAndCertificateRef = proto.String("/api/sslkeyandcertificate?name=" + *o.SslKeyAndCertificateRef)
		}
		if o.SslProfileRef != nil {
			o.SslProfileRef = proto.String("/api/sslprofile?name=" + *o.SslProfileRef)
		}
	}
}

