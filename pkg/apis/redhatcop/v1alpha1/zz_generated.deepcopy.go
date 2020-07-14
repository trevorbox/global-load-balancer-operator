// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	status "github.com/operator-framework/operator-sdk/pkg/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	out.CredentialsSecretRef = in.CredentialsSecretRef
	out.LoadBalancerServiceRef = in.LoadBalancerServiceRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDNSProviderConfig) DeepCopyInto(out *ExternalDNSProviderConfig) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDNSProviderConfig.
func (in *ExternalDNSProviderConfig) DeepCopy() *ExternalDNSProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalDNSProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDNSRecord) DeepCopyInto(out *GlobalDNSRecord) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDNSRecord.
func (in *GlobalDNSRecord) DeepCopy() *GlobalDNSRecord {
	if in == nil {
		return nil
	}
	out := new(GlobalDNSRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalDNSRecord) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDNSRecordList) DeepCopyInto(out *GlobalDNSRecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalDNSRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDNSRecordList.
func (in *GlobalDNSRecordList) DeepCopy() *GlobalDNSRecordList {
	if in == nil {
		return nil
	}
	out := new(GlobalDNSRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalDNSRecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDNSRecordSpec) DeepCopyInto(out *GlobalDNSRecordSpec) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		copy(*out, *in)
	}
	in.HealthCheck.DeepCopyInto(&out.HealthCheck)
	out.GlobalZoneRef = in.GlobalZoneRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDNSRecordSpec.
func (in *GlobalDNSRecordSpec) DeepCopy() *GlobalDNSRecordSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalDNSRecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDNSRecordStatus) DeepCopyInto(out *GlobalDNSRecordStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MonitoredServiceStatuses != nil {
		in, out := &in.MonitoredServiceStatuses, &out.MonitoredServiceStatuses
		*out = make(map[string]status.Conditions, len(*in))
		for key, val := range *in {
			var outVal []status.Condition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(status.Conditions, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.EndpointStatuses != nil {
		in, out := &in.EndpointStatuses, &out.EndpointStatuses
		*out = make(map[string]status.Conditions, len(*in))
		for key, val := range *in {
			var outVal []status.Condition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(status.Conditions, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDNSRecordStatus.
func (in *GlobalDNSRecordStatus) DeepCopy() *GlobalDNSRecordStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalDNSRecordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDNSZone) DeepCopyInto(out *GlobalDNSZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDNSZone.
func (in *GlobalDNSZone) DeepCopy() *GlobalDNSZone {
	if in == nil {
		return nil
	}
	out := new(GlobalDNSZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalDNSZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDNSZoneList) DeepCopyInto(out *GlobalDNSZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalDNSZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDNSZoneList.
func (in *GlobalDNSZoneList) DeepCopy() *GlobalDNSZoneList {
	if in == nil {
		return nil
	}
	out := new(GlobalDNSZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalDNSZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDNSZoneSpec) DeepCopyInto(out *GlobalDNSZoneSpec) {
	*out = *in
	in.Provider.DeepCopyInto(&out.Provider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDNSZoneSpec.
func (in *GlobalDNSZoneSpec) DeepCopy() *GlobalDNSZoneSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalDNSZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalDNSZoneStatus) DeepCopyInto(out *GlobalDNSZoneStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalDNSZoneStatus.
func (in *GlobalDNSZoneStatus) DeepCopy() *GlobalDNSZoneStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalDNSZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	if in.Route53 != nil {
		in, out := &in.Route53, &out.Route53
		*out = new(Route53ProviderConfig)
		**out = **in
	}
	if in.ExternalDNS != nil {
		in, out := &in.ExternalDNS, &out.ExternalDNS
		*out = new(ExternalDNSProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route53ProviderConfig) DeepCopyInto(out *Route53ProviderConfig) {
	*out = *in
	out.CredentialsSecretRef = in.CredentialsSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route53ProviderConfig.
func (in *Route53ProviderConfig) DeepCopy() *Route53ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(Route53ProviderConfig)
	in.DeepCopyInto(out)
	return out
}
