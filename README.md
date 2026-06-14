# Kakao Guard

Kakao Guard is a branded community fork of
[blurfx/KakaoTalkAdBlock](https://github.com/blurfx/KakaoTalkAdBlock), an
advertisement window blocker for the KakaoTalk Windows client.

## At a glance

![](https://raw.githubusercontent.com/blurfx/KakaoTalkAdBlock/main/kakaotalk.png)

This program runs in the background.

Right-click the tray icon to quit, or to check for new versions.

*Earlier versions of 2.2.0 can be quit by double-clicking the tray icon.*

![](https://raw.githubusercontent.com/blurfx/KakaoTalkAdBlock/main/tray.png)

## Startup

Use **Run on startup** from the tray menu to start Kakao Guard after signing
in to Windows. The registration follows the current executable path, so enable
it again after moving or renaming the executable.

If Windows previously disabled the app in Task Manager, enabling this option
again repairs that state.

When KakaoTalk is also configured to start at login, Kakao Guard safely
stores that command and launches KakaoTalk only after ad blocking is active.
Disabling **Run on startup** restores KakaoTalk's original startup registration.

Only one Kakao Guard process can run at a time.

## Windows Security

Release builds keep Go symbols and debugging metadata instead of stripping
them. This makes the executable easier to inspect and reduces false-positive
heuristics sometimes applied to unsigned Go GUI applications.

The application does not require a Microsoft Defender exclusion. Verify the
published SHA-256 hash before running a downloaded binary. A publicly trusted
code-signing certificate is still recommended for redistributed releases.
