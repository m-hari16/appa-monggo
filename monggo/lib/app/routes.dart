import 'package:get/get.dart';
import 'package:monggo/presentation/auth/login.dart';
import 'package:monggo/presentation/home/home.dart';

class Routes {
  static final List<GetPage> register = [
    GetPage(name: "/login", page: () => Login()),
    GetPage(name: "/", page: () => Home())
  ];
}
