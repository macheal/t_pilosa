// Code generated by "stringer -type=apiMethod"; DO NOT EDIT.

package pilosa

import "fmt"

const _apiMethod_name = "apiClusterMessageapiCreateFieldapiCreateFrameapiCreateIndexapiCreateInputDefinitionapiDeleteFieldapiDeleteFrameapiDeleteIndexapiDeleteInputDefinitionapiDeleteViewapiExportCSVapiFieldsapiFragmentBlockDataapiFragmentBlocksapiFrameAttrDiffapiImportapiImportValueapiIndexapiIndexAttrDiffapiInputDefinitionapiMarshalFragmentapiModifyFrameTimeQuantumapiQueryapiRecalculateCachesapiRemoveNodeapiResizeAbortapiRestoreFrameapiSetCoordinatorapiSliceNodesapiUnmarshalFragmentapiViewsapiWriteInput"

var _apiMethod_index = [...]uint16{0, 17, 31, 45, 59, 83, 97, 111, 125, 149, 162, 174, 183, 203, 220, 236, 245, 259, 267, 283, 301, 319, 344, 352, 372, 385, 399, 414, 431, 444, 464, 472, 485}

func (i apiMethod) String() string {
	if i < 0 || i >= apiMethod(len(_apiMethod_index)-1) {
		return fmt.Sprintf("apiMethod(%d)", i)
	}
	return _apiMethod_name[_apiMethod_index[i]:_apiMethod_index[i+1]]
}
