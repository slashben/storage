//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package softwarecomposition

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Annotation) DeepCopyInto(out *Annotation) {
	*out = *in
	out.Annotator = in.Annotator
	out.AnnotationSPDXIdentifier = in.AnnotationSPDXIdentifier
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Annotation.
func (in *Annotation) DeepCopy() *Annotation {
	if in == nil {
		return nil
	}
	out := new(Annotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Annotator) DeepCopyInto(out *Annotator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Annotator.
func (in *Annotator) DeepCopy() *Annotator {
	if in == nil {
		return nil
	}
	out := new(Annotator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArtifactOfProject) DeepCopyInto(out *ArtifactOfProject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArtifactOfProject.
func (in *ArtifactOfProject) DeepCopy() *ArtifactOfProject {
	if in == nil {
		return nil
	}
	out := new(ArtifactOfProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Checksum) DeepCopyInto(out *Checksum) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Checksum.
func (in *Checksum) DeepCopy() *Checksum {
	if in == nil {
		return nil
	}
	out := new(Checksum)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreationInfo) DeepCopyInto(out *CreationInfo) {
	*out = *in
	if in.Creators != nil {
		in, out := &in.Creators, &out.Creators
		*out = make([]Creator, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreationInfo.
func (in *CreationInfo) DeepCopy() *CreationInfo {
	if in == nil {
		return nil
	}
	out := new(CreationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Creator) DeepCopyInto(out *Creator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Creator.
func (in *Creator) DeepCopy() *Creator {
	if in == nil {
		return nil
	}
	out := new(Creator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocElementID) DeepCopyInto(out *DocElementID) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocElementID.
func (in *DocElementID) DeepCopy() *DocElementID {
	if in == nil {
		return nil
	}
	out := new(DocElementID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Document) DeepCopyInto(out *Document) {
	*out = *in
	if in.DocumentDescribes != nil {
		in, out := &in.DocumentDescribes, &out.DocumentDescribes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalDocumentReferences != nil {
		in, out := &in.ExternalDocumentReferences, &out.ExternalDocumentReferences
		*out = make([]ExternalDocumentRef, len(*in))
		copy(*out, *in)
	}
	if in.CreationInfo != nil {
		in, out := &in.CreationInfo, &out.CreationInfo
		*out = new(CreationInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]*Package, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Package)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]*File, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(File)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.OtherLicenses != nil {
		in, out := &in.OtherLicenses, &out.OtherLicenses
		*out = make([]*OtherLicense, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(OtherLicense)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Relationships != nil {
		in, out := &in.Relationships, &out.Relationships
		*out = make([]*Relationship, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Relationship)
				**out = **in
			}
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make([]Annotation, len(*in))
		copy(*out, *in)
	}
	if in.Snippets != nil {
		in, out := &in.Snippets, &out.Snippets
		*out = make([]Snippet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Reviews != nil {
		in, out := &in.Reviews, &out.Reviews
		*out = make([]*Review, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Review)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Document.
func (in *Document) DeepCopy() *Document {
	if in == nil {
		return nil
	}
	out := new(Document)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDocumentRef) DeepCopyInto(out *ExternalDocumentRef) {
	*out = *in
	out.Checksum = in.Checksum
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDocumentRef.
func (in *ExternalDocumentRef) DeepCopy() *ExternalDocumentRef {
	if in == nil {
		return nil
	}
	out := new(ExternalDocumentRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	if in.FileTypes != nil {
		in, out := &in.FileTypes, &out.FileTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Checksums != nil {
		in, out := &in.Checksums, &out.Checksums
		*out = make([]Checksum, len(*in))
		copy(*out, *in)
	}
	if in.LicenseInfoInFiles != nil {
		in, out := &in.LicenseInfoInFiles, &out.LicenseInfoInFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ArtifactOfProjects != nil {
		in, out := &in.ArtifactOfProjects, &out.ArtifactOfProjects
		*out = make([]*ArtifactOfProject, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ArtifactOfProject)
				**out = **in
			}
		}
	}
	if in.FileContributors != nil {
		in, out := &in.FileContributors, &out.FileContributors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FileAttributionTexts != nil {
		in, out := &in.FileAttributionTexts, &out.FileAttributionTexts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FileDependencies != nil {
		in, out := &in.FileDependencies, &out.FileDependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Snippets != nil {
		in, out := &in.Snippets, &out.Snippets
		*out = make(map[ElementID]*Snippet, len(*in))
		for key, val := range *in {
			var outVal *Snippet
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Snippet)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make([]Annotation, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Originator) DeepCopyInto(out *Originator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Originator.
func (in *Originator) DeepCopy() *Originator {
	if in == nil {
		return nil
	}
	out := new(Originator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OtherLicense) DeepCopyInto(out *OtherLicense) {
	*out = *in
	if in.LicenseCrossReferences != nil {
		in, out := &in.LicenseCrossReferences, &out.LicenseCrossReferences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OtherLicense.
func (in *OtherLicense) DeepCopy() *OtherLicense {
	if in == nil {
		return nil
	}
	out := new(OtherLicense)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Package) DeepCopyInto(out *Package) {
	*out = *in
	if in.HasFiles != nil {
		in, out := &in.HasFiles, &out.HasFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PackageSupplier != nil {
		in, out := &in.PackageSupplier, &out.PackageSupplier
		*out = new(Supplier)
		**out = **in
	}
	if in.PackageOriginator != nil {
		in, out := &in.PackageOriginator, &out.PackageOriginator
		*out = new(Originator)
		**out = **in
	}
	if in.PackageVerificationCode != nil {
		in, out := &in.PackageVerificationCode, &out.PackageVerificationCode
		*out = new(PackageVerificationCode)
		(*in).DeepCopyInto(*out)
	}
	if in.PackageChecksums != nil {
		in, out := &in.PackageChecksums, &out.PackageChecksums
		*out = make([]Checksum, len(*in))
		copy(*out, *in)
	}
	if in.PackageLicenseInfoFromFiles != nil {
		in, out := &in.PackageLicenseInfoFromFiles, &out.PackageLicenseInfoFromFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PackageExternalReferences != nil {
		in, out := &in.PackageExternalReferences, &out.PackageExternalReferences
		*out = make([]*PackageExternalReference, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PackageExternalReference)
				**out = **in
			}
		}
	}
	if in.PackageAttributionTexts != nil {
		in, out := &in.PackageAttributionTexts, &out.PackageAttributionTexts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]*File, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(File)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make([]Annotation, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Package.
func (in *Package) DeepCopy() *Package {
	if in == nil {
		return nil
	}
	out := new(Package)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageExternalReference) DeepCopyInto(out *PackageExternalReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageExternalReference.
func (in *PackageExternalReference) DeepCopy() *PackageExternalReference {
	if in == nil {
		return nil
	}
	out := new(PackageExternalReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageVerificationCode) DeepCopyInto(out *PackageVerificationCode) {
	*out = *in
	if in.ExcludedFiles != nil {
		in, out := &in.ExcludedFiles, &out.ExcludedFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageVerificationCode.
func (in *PackageVerificationCode) DeepCopy() *PackageVerificationCode {
	if in == nil {
		return nil
	}
	out := new(PackageVerificationCode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Relationship) DeepCopyInto(out *Relationship) {
	*out = *in
	out.RefA = in.RefA
	out.RefB = in.RefB
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Relationship.
func (in *Relationship) DeepCopy() *Relationship {
	if in == nil {
		return nil
	}
	out := new(Relationship)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportMeta) DeepCopyInto(out *ReportMeta) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportMeta.
func (in *ReportMeta) DeepCopy() *ReportMeta {
	if in == nil {
		return nil
	}
	out := new(ReportMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Review) DeepCopyInto(out *Review) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Review.
func (in *Review) DeepCopy() *Review {
	if in == nil {
		return nil
	}
	out := new(Review)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SBOMSPDXv2p3) DeepCopyInto(out *SBOMSPDXv2p3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SBOMSPDXv2p3.
func (in *SBOMSPDXv2p3) DeepCopy() *SBOMSPDXv2p3 {
	if in == nil {
		return nil
	}
	out := new(SBOMSPDXv2p3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SBOMSPDXv2p3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SBOMSPDXv2p3Filtered) DeepCopyInto(out *SBOMSPDXv2p3Filtered) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SBOMSPDXv2p3Filtered.
func (in *SBOMSPDXv2p3Filtered) DeepCopy() *SBOMSPDXv2p3Filtered {
	if in == nil {
		return nil
	}
	out := new(SBOMSPDXv2p3Filtered)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SBOMSPDXv2p3Filtered) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SBOMSPDXv2p3FilteredList) DeepCopyInto(out *SBOMSPDXv2p3FilteredList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SBOMSPDXv2p3Filtered, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SBOMSPDXv2p3FilteredList.
func (in *SBOMSPDXv2p3FilteredList) DeepCopy() *SBOMSPDXv2p3FilteredList {
	if in == nil {
		return nil
	}
	out := new(SBOMSPDXv2p3FilteredList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SBOMSPDXv2p3FilteredList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SBOMSPDXv2p3List) DeepCopyInto(out *SBOMSPDXv2p3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SBOMSPDXv2p3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SBOMSPDXv2p3List.
func (in *SBOMSPDXv2p3List) DeepCopy() *SBOMSPDXv2p3List {
	if in == nil {
		return nil
	}
	out := new(SBOMSPDXv2p3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SBOMSPDXv2p3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SBOMSPDXv2p3Spec) DeepCopyInto(out *SBOMSPDXv2p3Spec) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.SPDX.DeepCopyInto(&out.SPDX)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SBOMSPDXv2p3Spec.
func (in *SBOMSPDXv2p3Spec) DeepCopy() *SBOMSPDXv2p3Spec {
	if in == nil {
		return nil
	}
	out := new(SBOMSPDXv2p3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SBOMSPDXv2p3Status) DeepCopyInto(out *SBOMSPDXv2p3Status) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SBOMSPDXv2p3Status.
func (in *SBOMSPDXv2p3Status) DeepCopy() *SBOMSPDXv2p3Status {
	if in == nil {
		return nil
	}
	out := new(SBOMSPDXv2p3Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SPDXMeta) DeepCopyInto(out *SPDXMeta) {
	*out = *in
	out.Tool = in.Tool
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SPDXMeta.
func (in *SPDXMeta) DeepCopy() *SPDXMeta {
	if in == nil {
		return nil
	}
	out := new(SPDXMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Snippet) DeepCopyInto(out *Snippet) {
	*out = *in
	if in.Ranges != nil {
		in, out := &in.Ranges, &out.Ranges
		*out = make([]SnippetRange, len(*in))
		copy(*out, *in)
	}
	if in.LicenseInfoInSnippet != nil {
		in, out := &in.LicenseInfoInSnippet, &out.LicenseInfoInSnippet
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnippetAttributionTexts != nil {
		in, out := &in.SnippetAttributionTexts, &out.SnippetAttributionTexts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Snippet.
func (in *Snippet) DeepCopy() *Snippet {
	if in == nil {
		return nil
	}
	out := new(Snippet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnippetRange) DeepCopyInto(out *SnippetRange) {
	*out = *in
	out.StartPointer = in.StartPointer
	out.EndPointer = in.EndPointer
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnippetRange.
func (in *SnippetRange) DeepCopy() *SnippetRange {
	if in == nil {
		return nil
	}
	out := new(SnippetRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnippetRangePointer) DeepCopyInto(out *SnippetRangePointer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnippetRangePointer.
func (in *SnippetRangePointer) DeepCopy() *SnippetRangePointer {
	if in == nil {
		return nil
	}
	out := new(SnippetRangePointer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Supplier) DeepCopyInto(out *Supplier) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Supplier.
func (in *Supplier) DeepCopy() *Supplier {
	if in == nil {
		return nil
	}
	out := new(Supplier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolMeta) DeepCopyInto(out *ToolMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolMeta.
func (in *ToolMeta) DeepCopy() *ToolMeta {
	if in == nil {
		return nil
	}
	out := new(ToolMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityManifest) DeepCopyInto(out *VulnerabilityManifest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityManifest.
func (in *VulnerabilityManifest) DeepCopy() *VulnerabilityManifest {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityManifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityManifest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityManifestList) DeepCopyInto(out *VulnerabilityManifestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VulnerabilityManifest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityManifestList.
func (in *VulnerabilityManifestList) DeepCopy() *VulnerabilityManifestList {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityManifestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityManifestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityManifestSpec) DeepCopyInto(out *VulnerabilityManifestSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityManifestSpec.
func (in *VulnerabilityManifestSpec) DeepCopy() *VulnerabilityManifestSpec {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityManifestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityManifestStatus) DeepCopyInto(out *VulnerabilityManifestStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityManifestStatus.
func (in *VulnerabilityManifestStatus) DeepCopy() *VulnerabilityManifestStatus {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityManifestStatus)
	in.DeepCopyInto(out)
	return out
}
