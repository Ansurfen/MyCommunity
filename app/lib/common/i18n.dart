import 'package:flutter/cupertino.dart';
import 'package:get/get.dart';
import 'dart:ui' as ui;

class I18N extends Translations {
  @override
  Map<String, Map<String, String>> get keys => {
        'zh_CN': {'my_community': '社团云平台', 'splash_welcome': '欢迎使用社团云平台'},
        'en_US': {'my_community': 'My Community'}
      };
}

Locale getCurZone() {
  Locale curLocal = ui.window.locale;
  return curLocal.languageCode == 'zh'
      ? const Locale('zh', 'CN')
      : const Locale('en', 'US');
}

String? getFontFamilyByLanguage() {
  Locale curLocal = ui.window.locale;
  return curLocal.languageCode == 'zh' ? null : 'AverageSans';
}
