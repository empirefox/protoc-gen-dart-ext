import 'package:flutter/material.dart';
import 'package:pgde/pgde.dart' as $pgde;
import 'package:protobuf/protobuf.dart';

import 'config.pb.dart';
import 'config.l10n.dart';
import 'config.validate.dart';

class AdpRouterInputsBuilder extends $pgde.InputsBuilder {
  AdpRouterInputsBuilder($pgde.FormMessageData data)
      : super(data: data, oneofSize: 0);

  @override
  List<Widget> buildInputs(BuildContext context) {
    // TODO: implement buildInputs
    return null;
  }

  @override
  // TODO: implement hasError
  bool get hasError => null;

  @override
  // TODO: implement oneofBuilders
  Map<int, $pgde.InputsOneofBuilder> get oneofBuilders => null;

  @override
  // TODO: implement createEmptyDraft
  $pgde.CreateEmptyDraftFunc<GeneratedMessage> get createEmptyDraft => null;
}

class RouterItemInputsBuilder extends $pgde.InputsBuilder<RouterItem> {
  @override
  final $pgde.CreateEmptyDraftFunc createEmptyDraft;

  @override
  // sub-builder by tag
  final Map<int, $pgde.InputsOneofBuilder> oneofBuilders;

  RouterItemInputsBuilder($pgde.FormMessageData data)
      : createEmptyDraft = RouterItem.create,
        oneofBuilders = {
          3: AdpRouterInputsBuilder(data.childMessage(3)),
          4: $pgde.InputsMessageOneofBuilder(
              data: data.childMessage(4),
              createEmptyDraft: IPNetRouter.create,
              createValidator: (BuildContext context, $pgde.ValidateInfo vi) =>
                  IPNetRouterValidator(context, vi),
              createDecoration: (BuildContext context, String errorText) {
                final l10n = ConfigLocalizations.of(context);
                return InputDecoration(
                  labelText: l10n.IPNetRouterLabel,
                  hintText: l10n.IPNetRouterHint,
                  errorText: errorText,
                );
              }),
        },
        super(data: data, oneofSize: 1);

  @override
  bool get hasError => _hasErrorList.any((n) => n) || oneofHasError;
  final List<bool> _hasErrorList = List<bool>.filled(2, false);

  @override
  List<Widget> buildInputs(BuildContext context) {
    final l10n = ConfigLocalizations.of(context);
    final validator = RouterItemValidator(context, vi(context));

    oneofBuilders[5] = $pgde.InputsFieldOneofBuilder(
      decoration: InputDecoration(
        labelText: l10n.RouterItemNameLabel,
        hintText: l10n.RouterItemNameHint,
      ),
      validate: validator.assertField_name,
      build: (InputDecoration decoration, FormFieldValidator validator) =>
          TextFormField(
        decoration: decoration,
        initialValue: draft.name,
        onChanged: (n) => draft.name = n,
        autovalidate: true,
        validator: validator,
      ),
    );

    return <Widget>[
      TextFormField(
        decoration: InputDecoration(
          labelText: l10n.RouterItemNameLabel,
          hintText: l10n.RouterItemNameHint,
        ),
        initialValue: draft.name,
        onChanged: (n) => draft.name = n,
        autovalidate: true,
        validator: $pgde.InputsBuilder.catchError(
            validator.assertField_name, _hasErrorList, 0),
      ),
      ...buildOneofInputs(context, 0),
    ];
  }
}
