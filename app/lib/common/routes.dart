import 'package:get/get.dart';
import 'package:my_community/views/community/communityBinding.dart';
import 'package:my_community/views/community/communityPage.dart';
import 'package:my_community/views/home/homeBinding.dart';
import 'package:my_community/views/home/homePage.dart';
import 'package:my_community/views/login/loginBinding.dart';
import 'package:my_community/views/login/loginPage.dart';
import 'package:my_community/views/register/registerPage.dart';
import 'package:my_community/views/register/registerBinding.dart';
import 'package:my_community/views/splash/splashBiding.dart';
import 'package:my_community/views/splash/splashPage.dart';

abstract class Pages {
  static final pages = [
    GetPage(
        name: Routes.SPLASH,
        page: () => const SplashPage(),
        binding: SplashBinding()),
    GetPage(
        name: Routes.HOME,
        page: () => const HomePage(),
        binding: HomeBinding()),
    GetPage(
        name: Routes.LOGIN,
        page: () => const LoginPage(),
        binding: LoginBinding()),
    GetPage(
        name: Routes.REGISTER,
        page: () => const RegisterPage(),
        binding: RegisterBinding()),
    GetPage(
        name: Routes.MAIN,
        page: () => const CommunityPage(),
        binding: CommunityBinding())
  ];
}

abstract class Routes {
  static const HOME = '/';
  static const SPLASH = '/splash';
  static const ABOUT = '/about';
  static const BACKED = '/backed';
  static const LOGIN = '/login';
  static const REGISTER = '/register';
  static const MAIN = '/community';
}
