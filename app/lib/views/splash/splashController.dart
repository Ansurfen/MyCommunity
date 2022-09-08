import 'package:flutter/animation.dart';
import 'package:flutter/foundation.dart';
import 'package:get/get.dart';

class SplashController extends GetxController {
  String image = 'assets/splash.png';

  @override
  void onInit() {
    super.onInit();
    if (kDebugMode) {
      print("splash init");
    }
  }
}
