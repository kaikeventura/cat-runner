import 'dart:convert';
import 'dart:typed_data';

import '/flutter_flow/flutter_flow_util.dart';
import 'api_manager.dart';

export 'api_manager.dart' show ApiCallResponse;

const _kPrivateApiFunctionName = 'ffPrivateApiCall';

class HttpRunnerCall {
  static Future<ApiCallResponse> call({
    String? protocol = 'HTTP',
    String? host = '',
    int? port = 0,
    String? path = '',
    String? httpMethod = '',
    int? timeout,
    String? headers = '',
    String? parameters = '',
    String? contentText = '',
    int? usersAmount,
    int? interactionsAmount,
    int? interactionDelay,
  }) async {
    final ffApiRequestBody = '''
{
  "http": {
    "protocol": "${protocol}",
    "host": "${host}",
    "port": ${port},
    "path": "${path}",
    "httpMethod": "${httpMethod}",
    "timeout": ${timeout},
    "headers": ${headers},
    "parameters": ${parameters},
    "body": {
      "bodyFormat": "JSON",
      "contentText": "${contentText}"
    }
  },
  "virtualUser": {
    "usersAmount": ${usersAmount},
    "interactionsAmount": ${interactionsAmount},
    "interactionDelay": ${interactionDelay}
  }
}''';
    return ApiManager.instance.makeApiCall(
      callName: 'httpRunner',
      apiUrl: 'http://localhost:8080/v1/runner/http',
      callType: ApiCallType.POST,
      headers: {},
      params: {},
      body: ffApiRequestBody,
      bodyType: BodyType.JSON,
      returnBody: true,
      encodeBodyUtf8: false,
      decodeUtf8: false,
      cache: false,
    );
  }
}

class ApiPagingParams {
  int nextPageNumber = 0;
  int numItems = 0;
  dynamic lastResponse;

  ApiPagingParams({
    required this.nextPageNumber,
    required this.numItems,
    required this.lastResponse,
  });

  @override
  String toString() =>
      'PagingParams(nextPageNumber: $nextPageNumber, numItems: $numItems, lastResponse: $lastResponse,)';
}

String _serializeList(List? list) {
  list ??= <String>[];
  try {
    return json.encode(list);
  } catch (_) {
    return '[]';
  }
}

String _serializeJson(dynamic jsonVar, [bool isList = false]) {
  jsonVar ??= (isList ? [] : {});
  try {
    return json.encode(jsonVar);
  } catch (_) {
    return isList ? '[]' : '{}';
  }
}
