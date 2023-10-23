import '/backend/api_requests/api_calls.dart';
import '/flutter_flow/flutter_flow_drop_down.dart';
import '/flutter_flow/flutter_flow_theme.dart';
import '/flutter_flow/flutter_flow_util.dart';
import '/flutter_flow/flutter_flow_widgets.dart';
import '/flutter_flow/form_field_controller.dart';
import '/flutter_flow/custom_functions.dart' as functions;
import 'home_page_widget.dart' show HomePageWidget;
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

class HomePageModel extends FlutterFlowModel<HomePageWidget> {
  ///  State fields for stateful widgets in this page.

  final unfocusNode = FocusNode();
  // State field(s) for TabBar widget.
  TabController? tabBarController;
  int get tabBarCurrentIndex =>
      tabBarController != null ? tabBarController!.index : 0;

  // State field(s) for ProtocolDropDown widget.
  String? protocolDropDownValue;
  FormFieldController<String>? protocolDropDownValueController;
  // State field(s) for HostnameTextField widget.
  FocusNode? hostnameTextFieldFocusNode;
  TextEditingController? hostnameTextFieldController;
  String? Function(BuildContext, String?)? hostnameTextFieldControllerValidator;
  // State field(s) for PortTextField widget.
  FocusNode? portTextFieldFocusNode;
  TextEditingController? portTextFieldController;
  String? Function(BuildContext, String?)? portTextFieldControllerValidator;
  // State field(s) for HttpMethodDropDown widget.
  String? httpMethodDropDownValue;
  FormFieldController<String>? httpMethodDropDownValueController;
  // State field(s) for PathTextField widget.
  FocusNode? pathTextFieldFocusNode;
  TextEditingController? pathTextFieldController;
  String? Function(BuildContext, String?)? pathTextFieldControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue1;
  // State field(s) for QueryParamsTextField widget.
  FocusNode? queryParamsTextFieldFocusNode;
  TextEditingController? queryParamsTextFieldController;
  String? Function(BuildContext, String?)?
      queryParamsTextFieldControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue2;
  // State field(s) for HeaderParamsTextField widget.
  FocusNode? headerParamsTextFieldFocusNode;
  TextEditingController? headerParamsTextFieldController;
  String? Function(BuildContext, String?)?
      headerParamsTextFieldControllerValidator;
  // State field(s) for Switch widget.
  bool? switchValue3;
  // State field(s) for JsonBodyTextField widget.
  FocusNode? jsonBodyTextFieldFocusNode;
  TextEditingController? jsonBodyTextFieldController;
  String? Function(BuildContext, String?)? jsonBodyTextFieldControllerValidator;
  // State field(s) for UsersAmountTextField widget.
  FocusNode? usersAmountTextFieldFocusNode;
  TextEditingController? usersAmountTextFieldController;
  String? Function(BuildContext, String?)?
      usersAmountTextFieldControllerValidator;
  // State field(s) for InteractionsAmountTextField widget.
  FocusNode? interactionsAmountTextFieldFocusNode;
  TextEditingController? interactionsAmountTextFieldController;
  String? Function(BuildContext, String?)?
      interactionsAmountTextFieldControllerValidator;
  // State field(s) for InteractionDelayTextField widget.
  FocusNode? interactionDelayTextFieldFocusNode;
  TextEditingController? interactionDelayTextFieldController;
  String? Function(BuildContext, String?)?
      interactionDelayTextFieldControllerValidator;
  // Stores action output result for [Backend Call - API (httpRunner)] action in Button widget.
  ApiCallResponse? httpRunnerResponse;

  /// Initialization and disposal methods.

  void initState(BuildContext context) {}

  void dispose() {
    unfocusNode.dispose();
    tabBarController?.dispose();
    hostnameTextFieldFocusNode?.dispose();
    hostnameTextFieldController?.dispose();

    portTextFieldFocusNode?.dispose();
    portTextFieldController?.dispose();

    pathTextFieldFocusNode?.dispose();
    pathTextFieldController?.dispose();

    queryParamsTextFieldFocusNode?.dispose();
    queryParamsTextFieldController?.dispose();

    headerParamsTextFieldFocusNode?.dispose();
    headerParamsTextFieldController?.dispose();

    jsonBodyTextFieldFocusNode?.dispose();
    jsonBodyTextFieldController?.dispose();

    usersAmountTextFieldFocusNode?.dispose();
    usersAmountTextFieldController?.dispose();

    interactionsAmountTextFieldFocusNode?.dispose();
    interactionsAmountTextFieldController?.dispose();

    interactionDelayTextFieldFocusNode?.dispose();
    interactionDelayTextFieldController?.dispose();
  }

  /// Action blocks are added here.

  /// Additional helper methods are added here.
}
