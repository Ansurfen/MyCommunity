import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:my_community/common/routes.dart';
import 'package:my_community/views/splash/splashController.dart';

class SplashPage extends StatelessWidget {
  const SplashPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
        backgroundColor: Colors.white, body: _SplashPageWidget());
  }
}

class _SplashPageWidget extends StatefulWidget {
  const _SplashPageWidget({Key? key}) : super(key: key);

  @override
  State<_SplashPageWidget> createState() => _SplashPageWidgetState();
}

class _SplashPageWidgetState extends State<_SplashPageWidget> {
  SplashController splashController = Get.find<SplashController>();


  @override
  Widget build(BuildContext context) {
    return Stack(
      fit: StackFit.expand,
      alignment: Alignment.center,
      children: [
        Container(
          decoration: const BoxDecoration(
            color: Colors.purple,
          ),
          child: Image.asset(splashController.image),
        ),
        Container(
          margin: const EdgeInsets.only(top: 100),
          child: Text(
            "社团云平台".tr,
            textAlign: TextAlign.center,
            style: const TextStyle(
              fontSize: 21,
              fontWeight: FontWeight.w300,
              color: Colors.white
            ),
          ),
        )
      ],
    );
  }

  @override
  void initState() {
    super.initState();
    var duration = const Duration(seconds: 1);
    Future.delayed(duration, goToHomePage);
  }

  void goToHomePage() {
    Get.offNamed(Routes.HOME);
  }
}
