import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_community/views/community/communityController.dart';

class CommunityPage extends StatefulWidget {
  const CommunityPage({Key? key}) : super(key: key);

  @override
  State<CommunityPage> createState() => _CommunityPageState();
}

class _CommunityPageState extends State<CommunityPage> {
  CommunityController communityController = Get.find<CommunityController>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Text(Get.parameters['id']!),
    );
  }
}
