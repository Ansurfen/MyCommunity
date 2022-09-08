import 'package:get/get.dart';
import 'package:my_community/views/home/homeController.dart';

class HomeBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<HomeController>(HomeController());
  }
}