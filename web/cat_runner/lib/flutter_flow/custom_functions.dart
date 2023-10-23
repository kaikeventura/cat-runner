import 'dart:convert';
import 'dart:math' as math;

import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:intl/intl.dart';
import 'package:timeago/timeago.dart' as timeago;
import 'lat_lng.dart';
import 'place.dart';
import 'uploaded_file.dart';
import '/backend/schema/structs/index.dart';

String keyValueFunction(String? queryParams) {
  if (queryParams == null) return '[]';

  List<Map<String, String>> objects = [];

  List<String> pairs = queryParams.split(';');
  for (String pair in pairs) {
    List<String> keyValue = pair.split(':');
    if (keyValue.length == 2) {
      String key = keyValue[0].trim();
      String value = keyValue[1].trim();
      Map<String, String> object = {
        "key": key,
        "value": value,
      };
      objects.add(object);
    }
  }

  return jsonEncode(objects);
}

String? jsonStringTransformFunction(String? json) {
  return json != null ? json.replaceAll('"', r'\"') : null;
}
