# fdroid
This repository hosts a [F-Droid](https://f-droid.org/) repo for my apps. This repository allows you to get the latest and greatest app updates faster than the app stores. 

### Apps

<!-- This table is auto-generated. Do not edit -->
| Icon | Name | Description | Version |
| --- | --- | --- | --- |
| <a href="https://github.com/wgtunnel/wgtunnel"><img src="fdroid/repo/com.zaneschepke.wireguardautotunnel/en-US/icon.png" alt="WG Tunnel icon" width="36px" height="36px"></a> | [**WG Tunnel**](https://github.com/wgtunnel/wgtunnel) | A FOSS Android client for WireGuard and AmneziaWG with auto-tunneling. | 3.9.0 (38900) |
<!-- end apps table -->

### How to use
1. First, you should [install the F-Droid app](https://f-droid.org/), it's an alternative app store for Android.
2. In the F-droid app, navigate to Settings > Repositories.
3. Click the + floating action button to add a new repository. 
4. Click *SCAN QR CODE* and scan the QR below:

    <p align="center">
      <img src=".github/qrcode.png?raw=true" alt="F-Droid repo QR code" style="width:500px;height:500px;"/>
    </p>

    Alternatively, you can also add the repository manually with the following url:

      ```
    https://raw.githubusercontent.com/zaneschepke/fdroid/main/fdroid/repo?fingerprint=0890C5D44C0109E366801C39840325E810E21B270B9D2AEC53CE0D6C5FC849DB
    ```

4. You can now install my apps, e.g. start by searching for "WG Tunnel" in the F-Droid client.

Please note that some apps published here might contain [Anti-Features](https://f-droid.org/en/docs/Anti-Features/). If you can't find an app by searching for it, you can go to settings and enable "Include anti-feature apps".

### For developers
If you are a developer and want to publish your own apps right from GitHub Actions as an F-Droid repo, you can fork/copy this repo and see  [the documentation](setup.md) for more information on how to set it up.

### [License](LICENSE)
The license is for the files in this repository, *except* those in the `fdroid` directory. These files *might* be licensed differently; you can use an F-Droid client to get the details for each app.
