import 'package:get/get.dart';
import 'package:my_community/views/register/registerController.dart';

class RegisterBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<RegisterController>(RegisterController());
  }
}