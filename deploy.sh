~/go/bin/qtdeploy -docker build windows_32_static
~/go/bin/qtdeploy build
cd deploy
cp -ru linux/* mycrypto-qt/
cp -ru windows/* mycrypto-qt-win32/
rm mycrypto-qt.tar.xz mycrypto-qt-win32.zip
tar cvfJ mycrypto-qt.tar.xz mycrypto-qt
zip -r9 mycrypto-qt-win32.zip mycrypto-qt-win32

