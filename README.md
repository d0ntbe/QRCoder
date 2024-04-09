# QRCoder
QR generator for Mobile Security testing. 
For QR gen uses this "github.com/skip2/go-qrcode". Thx! 


USAGE:

Example: go run main.go deeplink://aplication/api/v1/ppp

Example: go run main.go deeplink://aplication/api/v1/ppp?title=hello#popla=gopla#index=123#toto=3v

Example: QRCoder.exe superbank://linktoapp/p2p

If use Windows - do not forgot double quotes. Example: QRCoder.exe "ddd://ooooo?dsdsasd=123&dddaq=1"

Payloads are in payloads.txt. But you can also add your specified payloads, line by line.


![image](https://github.com/d0ntbe/QRCoder/assets/88555610/d6c7e865-2216-46bc-b04c-ae28a22e9730)

Vulns: 

1) Open Redirect
2) XSS
3) DoS
4) Fishing
5) Banking: payment by details from
6) Local SQLi (Android, IOS)

Examples:
 
appbank://link?https://www.p@yment.com/account=78127312936781&sum=100&cur=rub -> ![image](https://github.com/d0ntbe/QRCoder/assets/88555610/7bd23cb9-9108-456f-8d18-53e428cd8180)

appbank://link?https://www.p@yment.com/account=youtubik.ev11il.com&sum=100&cur=rub -> ![image](https://github.com/d0ntbe/QRCoder/assets/88555610/d566f0b2-a34f-4caf-b469-d07ee9f69f3d)


