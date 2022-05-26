# Bridge Problem

In the example, we have 2 types of devices - **Radio** and **TV**. We also have 2 **RemoteControl** classes which use concrete devices and are fully dependent on it. If there is a small change in a device class, there will be a small change in a remote class. Therefore, we need to get rid of this dependency.

Also, imagine if we need to add a new device (e.g. **MusicPlayer**), we'd need to add a new device class and a new remote class as well, what might not be convenient to do.
