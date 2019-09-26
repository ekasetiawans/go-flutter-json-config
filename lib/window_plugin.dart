import 'package:flutter/services.dart';

class WindowPlugin {
  static MethodChannel channel =
      MethodChannel("com.github.ekasetiawans/window_plugin");

  static setFullScreen(bool value) async {
    await channel.invokeMethod("setFullScreen", value);
  }

  static Future<bool> isFullScreen() async {
    return await channel.invokeMethod("getFullScreen");
  }
}
