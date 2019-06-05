// Code generated by "stringer -type=apiMethod"; DO NOT EDIT.

package pilosa

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[apiClusterMessage-0]
	_ = x[apiCreateField-1]
	_ = x[apiCreateIndex-2]
	_ = x[apiDeleteField-3]
	_ = x[apiDeleteAvailableShard-4]
	_ = x[apiDeleteIndex-5]
	_ = x[apiDeleteView-6]
	_ = x[apiExportCSV-7]
	_ = x[apiFragmentBlockData-8]
	_ = x[apiFragmentBlocks-9]
	_ = x[apiFragmentData-10]
	_ = x[apiField-11]
	_ = x[apiFieldAttrDiff-12]
	_ = x[apiImport-13]
	_ = x[apiImportValue-14]
	_ = x[apiIndex-15]
	_ = x[apiIndexAttrDiff-16]
	_ = x[apiQuery-17]
	_ = x[apiRecalculateCaches-18]
	_ = x[apiRemoveNode-19]
	_ = x[apiResizeAbort-20]
	_ = x[apiSetCoordinator-21]
	_ = x[apiShardNodes-22]
	_ = x[apiViews-23]
	_ = x[apiApplySchema-24]
}

const _apiMethod_name = "apiClusterMessageapiCreateFieldapiCreateIndexapiDeleteFieldapiDeleteAvailableShardapiDeleteIndexapiDeleteViewapiExportCSVapiFragmentBlockDataapiFragmentBlocksapiFragmentDataapiFieldapiFieldAttrDiffapiImportapiImportValueapiIndexapiIndexAttrDiffapiQueryapiRecalculateCachesapiRemoveNodeapiResizeAbortapiSetCoordinatorapiShardNodesapiViewsapiApplySchema"

var _apiMethod_index = [...]uint16{0, 17, 31, 45, 59, 82, 96, 109, 121, 141, 158, 173, 181, 197, 206, 220, 228, 244, 252, 272, 285, 299, 316, 329, 337, 351}

func (i apiMethod) String() string {
	if i < 0 || i >= apiMethod(len(_apiMethod_index)-1) {
		return "apiMethod(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _apiMethod_name[_apiMethod_index[i]:_apiMethod_index[i+1]]
}
