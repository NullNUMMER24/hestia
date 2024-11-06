// lib/main.dart
import 'package:flutter/material.dart';
import 'dashboard.dart';  // Import the second page
//import 'login.dart';   // Import the third page

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter startsite',
      home: HomePage(),
      theme: ThemeData(
	scaffoldBackgroundColor: const Color(0xFFEFEFEF),
      )
    );
  }
}

class HomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Home Page'),
      ),
      drawer: Drawer(
        child: ListView(
          padding: EdgeInsets.zero,
          children: <Widget>[
            DrawerHeader(
              child: Text(
                'Menu',
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 24,
                ),
              ),
              decoration: BoxDecoration(
                color: Colors.blue,
              ),
            Image.asset( //Debug - remove this
	     'assets/hestia_logo.png',
	     height: 200,
	    ),
            ),
            ListTile(
              title: Text('Dashboard'),
              onTap: () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => dashboard()),
                );
              },
            ),
           // ListTile(
           //   title: Text('Login'),
           //   onTap: () {
           //     Navigator.push(
           //       context,
           //       MaterialPageRoute(builder: (context) => login()),
           //     );
           //   },
           // ),
          ],
        ),
      ),
      body: Center(
        child: Text('Welcome to the Home Page!'),
      ),
    );
  }
}
