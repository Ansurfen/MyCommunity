import 'package:get/get.dart';
import 'package:my_community/views/login/loginController.dart';

class LoginBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<LoginController>(LoginController());
  }
}