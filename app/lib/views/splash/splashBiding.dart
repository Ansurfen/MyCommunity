import 'package:get/get.dart';
import 'package:my_community/views/splash/splashController.dart';

class SplashBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<SplashController>(SplashController());
  }
}