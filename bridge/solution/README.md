# Bridge Solution

So, we unified our remote classes into a single one - **RemoteControl**, which accepts the device in the constructor and works with it through the **Device** interface. Now the remote class code doesn't depend on a concrete device, and we can add as many devices as we want. Changes in device class wouldn't change the **RemoteControl** class. The _device_ property in this class is the bridge itself.
