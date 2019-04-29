///
//  Generated code. Do not modify.
//  source: protos/form/form.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

const UnitPlace$json = const {
  '1': 'UnitPlace',
  '2': const [
    const {'1': 'noUnitPlace', '2': 0},
    const {'1': 'suffix', '2': 1},
    const {'1': 'labelSuffix', '2': 2},
  ],
};

const InputOption$json = const {
  '1': 'InputOption',
  '2': const [
    const {'1': 'icon', '3': 1, '4': 1, '5': 14, '6': '.form.MaterialIcon', '10': 'icon'},
  ],
};

const Field$json = const {
  '1': 'Field',
  '2': const [
    const {'1': 'decoration', '3': 1, '4': 1, '5': 11, '6': '.form.InputDecoration', '10': 'decoration'},
    const {'1': 'hide', '3': 2, '4': 1, '5': 8, '10': 'hide'},
    const {'1': 'enabled', '3': 3, '4': 1, '5': 8, '10': 'enabled'},
    const {'1': 'ensureAgain', '3': 4, '4': 1, '5': 8, '10': 'ensureAgain'},
    const {'1': 'emptyAsNoChange', '3': 5, '4': 1, '5': 8, '10': 'emptyAsNoChange'},
    const {'1': 'disableDesc', '3': 6, '4': 1, '5': 8, '10': 'disableDesc'},
    const {'1': 'text', '3': 11, '4': 1, '5': 11, '6': '.form.TextInput', '9': 0, '10': 'text'},
    const {'1': 'typeAhead', '3': 12, '4': 1, '5': 11, '6': '.form.TypeAheadInput', '9': 0, '10': 'typeAhead'},
    const {'1': 'signature', '3': 13, '4': 1, '5': 11, '6': '.form.SignatureInput', '9': 0, '10': 'signature'},
    const {'1': 'number', '3': 14, '4': 1, '5': 11, '6': '.form.NumberInput', '9': 0, '10': 'number'},
    const {'1': 'stepper', '3': 15, '4': 1, '5': 11, '6': '.form.StepperInput', '9': 0, '10': 'stepper'},
    const {'1': 'slider', '3': 16, '4': 1, '5': 11, '6': '.form.SliderInput', '9': 0, '10': 'slider'},
    const {'1': 'rate', '3': 17, '4': 1, '5': 11, '6': '.form.RateInput', '9': 0, '10': 'rate'},
    const {'1': 'radio', '3': 18, '4': 1, '5': 11, '6': '.form.RadioInput', '9': 0, '10': 'radio'},
    const {'1': 'boolSwitch', '3': 19, '4': 1, '5': 11, '6': '.form.SwitchInput', '9': 0, '10': 'boolSwitch'},
    const {'1': 'list', '3': 20, '4': 1, '5': 11, '6': '.form.ListInput', '9': 0, '10': 'list'},
    const {'1': 'datetime', '3': 21, '4': 1, '5': 11, '6': '.form.DatetimePickerInput', '9': 0, '10': 'datetime'},
  ],
  '8': const [
    const {'1': 'type'},
  ],
};

const TextInput$json = const {
  '1': 'TextInput',
  '2': const [
    const {'1': 'keyboardType', '3': 1, '4': 1, '5': 14, '6': '.form.TextInput.Type', '10': 'keyboardType'},
    const {'1': 'obscureText', '3': 2, '4': 1, '5': 8, '10': 'obscureText'},
    const {'1': 'autocorrect', '3': 3, '4': 1, '5': 8, '10': 'autocorrect'},
    const {'1': 'maxLines', '3': 4, '4': 1, '5': 5, '10': 'maxLines'},
    const {'1': 'maxLengthEnforced', '3': 5, '4': 1, '5': 8, '10': 'maxLengthEnforced'},
  ],
  '4': const [TextInput_Type$json],
};

const TextInput_Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'autoByType', '2': 0},
    const {'1': 'datetime', '2': 2},
    const {'1': 'emailAddress', '2': 3},
    const {'1': 'multiline', '2': 4},
    const {'1': 'number', '2': 5},
    const {'1': 'phone', '2': 6},
    const {'1': 'text', '2': 7},
    const {'1': 'url', '2': 8},
  ],
};

const TypeAheadInput$json = const {
  '1': 'TypeAheadInput',
  '2': const [
    const {'1': 'codeAccessor', '3': 1, '4': 1, '5': 9, '9': 0, '10': 'codeAccessor'},
    const {'1': 'fromEnum', '3': 2, '4': 1, '5': 8, '9': 0, '10': 'fromEnum'},
  ],
  '8': const [
    const {'1': 'builder'},
  ],
};

const SignatureInput$json = const {
  '1': 'SignatureInput',
  '2': const [
    const {'1': 'width', '3': 1, '4': 1, '5': 1, '10': 'width'},
    const {'1': 'height', '3': 2, '4': 1, '5': 1, '10': 'height'},
    const {'1': 'backgroundColor', '3': 3, '4': 1, '5': 11, '6': '.form.MaterialColor', '10': 'backgroundColor'},
    const {'1': 'penColor', '3': 4, '4': 1, '5': 11, '6': '.form.MaterialColor', '10': 'penColor'},
    const {'1': 'penStrokeWidth', '3': 5, '4': 1, '5': 1, '10': 'penStrokeWidth'},
  ],
};

const NumberInput$json = const {
  '1': 'NumberInput',
  '2': const [
    const {'1': 'fullUnit', '3': 1, '4': 1, '5': 8, '10': 'fullUnit'},
    const {'1': 'unitPlace', '3': 2, '4': 1, '5': 14, '6': '.form.UnitPlace', '10': 'unitPlace'},
  ],
};

const StepperInput$json = const {
  '1': 'StepperInput',
  '2': const [
    const {'1': 'fullUnit', '3': 1, '4': 1, '5': 8, '10': 'fullUnit'},
    const {'1': 'unitPlace', '3': 2, '4': 1, '5': 14, '6': '.form.UnitPlace', '10': 'unitPlace'},
    const {'1': 'step', '3': 3, '4': 1, '5': 13, '10': 'step'},
  ],
};

const SliderInput$json = const {
  '1': 'SliderInput',
  '2': const [
    const {'1': 'fullUnit', '3': 1, '4': 1, '5': 8, '10': 'fullUnit'},
    const {'1': 'unitPlace', '3': 2, '4': 1, '5': 14, '6': '.form.UnitPlace', '10': 'unitPlace'},
    const {'1': 'divisions', '3': 3, '4': 1, '5': 13, '10': 'divisions'},
  ],
};

const RateInput$json = const {
  '1': 'RateInput',
  '2': const [
    const {'1': 'icon', '3': 1, '4': 1, '5': 14, '6': '.form.MaterialIcon', '10': 'icon'},
  ],
};

const RadioInput$json = const {
  '1': 'RadioInput',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.form.RadioInput.Type', '10': 'type'},
    const {'1': 'codeAccessor', '3': 2, '4': 1, '5': 9, '9': 0, '10': 'codeAccessor'},
    const {'1': 'fromEnum', '3': 3, '4': 1, '5': 8, '9': 0, '10': 'fromEnum'},
  ],
  '4': const [RadioInput_Type$json],
  '8': const [
    const {'1': 'builder'},
  ],
};

const RadioInput_Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'radio', '2': 0},
    const {'1': 'dropdown', '2': 1},
    const {'1': 'segmentedControl', '2': 2},
  ],
};

const SwitchInput$json = const {
  '1': 'SwitchInput',
  '2': const [
    const {'1': 'checkbox', '3': 1, '4': 1, '5': 8, '10': 'checkbox'},
  ],
};

const ListInput$json = const {
  '1': 'ListInput',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.form.ListInput.Type', '10': 'type'},
    const {'1': 'codeAccessor', '3': 2, '4': 1, '5': 9, '9': 0, '10': 'codeAccessor'},
    const {'1': 'fromEnum', '3': 3, '4': 1, '5': 8, '9': 0, '10': 'fromEnum'},
  ],
  '4': const [ListInput_Type$json],
  '8': const [
    const {'1': 'builder'},
  ],
};

const ListInput_Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'checkbox', '2': 0},
    const {'1': 'chips', '2': 1},
  ],
};

const DatetimePickerInput$json = const {
  '1': 'DatetimePickerInput',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.form.DatetimePickerInput.Type', '10': 'type'},
    const {'1': 'format', '3': 2, '4': 1, '5': 9, '10': 'format'},
  ],
  '4': const [DatetimePickerInput_Type$json],
};

const DatetimePickerInput_Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'date', '2': 0},
    const {'1': 'time', '2': 1},
    const {'1': 'both', '2': 2},
  ],
};

const InputDecoration$json = const {
  '1': 'InputDecoration',
  '2': const [
    const {'1': 'icon', '3': 1, '4': 1, '5': 14, '6': '.form.MaterialIcon', '10': 'icon'},
    const {'1': 'prefixIcon', '3': 2, '4': 1, '5': 14, '6': '.form.MaterialIcon', '10': 'prefixIcon'},
    const {'1': 'suffixIcon', '3': 3, '4': 1, '5': 14, '6': '.form.MaterialIcon', '10': 'suffixIcon'},
    const {'1': 'hasCounter', '3': 4, '4': 1, '5': 8, '10': 'hasCounter'},
  ],
};

const Date$json = const {
  '1': 'Date',
  '2': const [
    const {'1': 'unix', '3': 1, '4': 1, '5': 3, '9': 0, '10': 'unix'},
    const {'1': 'rfc3339', '3': 2, '4': 1, '5': 9, '9': 0, '10': 'rfc3339'},
  ],
  '8': const [
    const {'1': 'utc'},
  ],
};

