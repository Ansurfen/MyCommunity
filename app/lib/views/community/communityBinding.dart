import 'package:get/get.dart';
import 'package:my_community/views/community/communityController.dart';

class CommunityBinding extends Bindings {
  @override
  void dependencies() {
    Get.put<CommunityController>(CommunityController());
  }
}