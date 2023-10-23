// ignore_for_file: unnecessary_getters_setters

import '/backend/schema/util/schema_util.dart';

import 'index.dart';
import '/flutter_flow/flutter_flow_util.dart';

class QueryParamsStruct extends BaseStruct {
  QueryParamsStruct({
    List<QueryParamsStruct>? queryParams,
  }) : _queryParams = queryParams;

  // "queryParams" field.
  List<QueryParamsStruct>? _queryParams;
  List<QueryParamsStruct> get queryParams => _queryParams ?? const [];
  set queryParams(List<QueryParamsStruct>? val) => _queryParams = val;
  void updateQueryParams(Function(List<QueryParamsStruct>) updateFn) =>
      updateFn(_queryParams ??= []);
  bool hasQueryParams() => _queryParams != null;

  static QueryParamsStruct fromMap(Map<String, dynamic> data) =>
      QueryParamsStruct(
        queryParams: getStructList(
          data['queryParams'],
          QueryParamsStruct.fromMap,
        ),
      );

  static QueryParamsStruct? maybeFromMap(dynamic data) =>
      data is Map<String, dynamic> ? QueryParamsStruct.fromMap(data) : null;

  Map<String, dynamic> toMap() => {
        'queryParams': _queryParams?.map((e) => e.toMap()).toList(),
      }.withoutNulls;

  @override
  Map<String, dynamic> toSerializableMap() => {
        'queryParams': serializeParam(
          _queryParams,
          ParamType.DataStruct,
          true,
        ),
      }.withoutNulls;

  static QueryParamsStruct fromSerializableMap(Map<String, dynamic> data) =>
      QueryParamsStruct(
        queryParams: deserializeStructParam<QueryParamsStruct>(
          data['queryParams'],
          ParamType.DataStruct,
          true,
          structBuilder: QueryParamsStruct.fromSerializableMap,
        ),
      );

  @override
  String toString() => 'QueryParamsStruct(${toMap()})';

  @override
  bool operator ==(Object other) {
    const listEquality = ListEquality();
    return other is QueryParamsStruct &&
        listEquality.equals(queryParams, other.queryParams);
  }

  @override
  int get hashCode => const ListEquality().hash([queryParams]);
}

QueryParamsStruct createQueryParamsStruct() => QueryParamsStruct();
