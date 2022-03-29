#include <Arduino.h>
//#include <M5stack.h>

void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
  Serial.print("やあ");
//  M5.begin();
}

void loop() {
  Serial.println("やあ、アリアルさんだよ。");
 // M5.Lcd.drawJpgFile(SD, "/arialsou1.jpg");
  // put your main code here, to run repeatedly:
}