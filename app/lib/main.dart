import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';
import 'package:my_community/common/i18n.dart';
import 'package:my_community/common/routes.dart';
import 'package:oktoast/oktoast.dart';

void main() {
  runApp(createApp());
}

Widget createApp() {
  return OKToast(
    dismissOtherOnShow: true,
    child: GetMaterialApp(
      debugShowCheckedModeBanner: false,
      translations: I18N(),
      locale: getCurZone(),
      fallbackLocale: const Locale('en','US'),
      builder: EasyLoading.init(),
      theme: ThemeData(
        primaryColor: Colors.white,
        fontFamily: getFontFamilyByLanguage()
      ),
      initialRoute: Routes.SPLASH,
      getPages: Pages.pages,
    ),
  );
}