import 'package:flutter/material.dart';
import 'package:json_config/window_plugin.dart';

class HomePage extends StatefulWidget {
  HomePage({Key key}) : super(key: key);

  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Window Plugin"),
      ),
      body: FutureBuilder(
        future: WindowPlugin.isFullScreen(),
        builder: (context, snapshot) {
          return Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: <Widget>[
                Switch(
                  value: snapshot.data ?? false,
                  onChanged: (value) async {
                    await WindowPlugin.setFullScreen(value);
                    Scaffold.of(context).showSnackBar(
                      SnackBar(
                        content: Text("Saved. Please restart application."),
                      ),
                    );
                    setState(() {});
                  },
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
