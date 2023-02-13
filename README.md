# 永久存储合约

[![Smart Contract](https://badgen.net/badge/smart-contract/Solidity/orange)](https://soliditylang.org/) [![License](https://badgen.net/badge/license/MIT/blue)](https://typescriptlang.org)

数据存储关乎任何一个IT系统，站在操作类型上说数据存储，可归纳为增删改查。随着人们思想认知的不断提升推动了IT的技术更新，时至今日，更多IT系统内的增删改查均已融入了多技术，区块链便是其中之一。区块链是一种技术集，虽然也在持续优化更新，但依旧有其优势和劣势。

任何一个区块链应用，都离不开与原有的传统IT系统进行交互，可以说传统IT系统是区块链应用的“执行者”。所以，利用好区块链技术可以对传统IT系统进行优化，甚至是解决原有不可解决的一些问题；同时，传统IT系统可以补充区块链技术的短板，两者优势互补必然出现更优的结果。

## 1. 使用说明 
* 合约内引用了 Openzeppelin 的 Ownable 库，所以在合约方法上添加 onlyOwner 修饰器，可确保合约方法仅支持合约拥有者（即部署合约的人）调用。

* 随着业务系统的迭代更新，合约完全有可能也会涉及到更新，合约的更新必须保证不影响旧数据，所以合约代码在首次设计的时候，一定要要做好逻辑代码和数据代码的剥离。

* 和传统IT系统类似，合约需要支持升级。合约升级有多种方案可以实现，使用较多的是采用透明代理和 UUPS，本合约采用的是 UUPS 模式，所以合约组成结构上多了一个代理合约，用户在业务系统只需调用代理合约即可。需要注意，无论是哪一种升级模式，在部署或者升级合约过程中都需先部署主合约，成功后才可部署代理合约。

* 合约的安全问题不容小觑，所以合约内引用了 Openzeppelin 的 Pausable 库，如发现合约漏洞或缺陷，用户可以暂停所有人继续调用合约。

## 2. 部署及升级说明

**1）部署**

第一步：部署主合约。

第二步：部署代理合约。
部署时传参主合约地址、initialize的方法签名（“0x8129fc1c”），实现与主合约的映射及初始化操作。

注意：用Java Web3J部署代理合约时，要注意在部署的构造函数入参中，要把合约地址和initialize方法签名（“0x8129fc1c”）去掉0x填充 0 拼接在合约bin后面来部署。用Remix部署代理合约时构造函数入参中合约地址和initialize方法签名0x要保留。

**2）升级**

第一步：部署新版本的主合约。

第二步：调用代理合约中的upgradeTo方法，参数为新版本的主合约地址，实现与新版本主合约的映射，达到升级目的。

## 3.合约概述
用户部署合约后，可以通过调用合约方法的方式对合约进行暂停（pause）、启动（unpause）和升级（upgradeTo）操作，同时可以将业务数据提交到链上进行存储（storeData），并对其进行删除（deleteData）、修改（storeData）、查询（queryData）操作。

* 数据上链：业务系统内自定义 key 值，将业务数据作为 value，调用 storeData 方法可完成数据上链。
* 删除：业务系统内指定需要删除的业务数据的 key 值，调用 deleteData 方法可完成数据的删除。
* 修改：业务系统内指定需要修改的业务数据的 key 值，将新的业务数据作为 value，调用 storeData 方法可完成数据的修改。
* 查询：业务系统内指定需要查询的业务数据的 key 值，调用 queryData 方法可查到链上存储的业务数据。


## 4.合约方法

编号  |  方法声明   |  方法描述 
--------|------------|------------
&nbsp;1 | *function storeData(string memory key,string memory value) public onlyOwner whenNotPaused* |      key-value 形式存储、修改业务数据              
&nbsp;2 | *function deleteData(string memory key) public onlyOwner whenNotPaused* | 根据 key 删除对应的业务数据
&nbsp;3 | *function queryData(string memory key) public view whenNotPaused returns (string memory)* |      根据 key 获取对应的业务数据  
            


## 5.Go语言调用示例

### 1）数据上链
```
func TestStoreData(t *testing.T) {
    cli, err := ethclient.Dial(NodeUrl)
    if err != nil {
        log.Logger.Error(err)
    }
    instance, err := storage.NewStorage(common.HexToAddress(Address), cli)
    if err != nil {
        log.Logger.Error(err)
    }
    auth, err := eth.GenAuth(cli, PrivateKey)
    if err != nil {
        log.Logger.Error(err)
    }
    tx, err := instance.StoreData(auth, key1, value1)
    if err != nil {
        log.Logger.Error(err)
    }
    fmt.Println("tx Hash:", tx.Hash().String())
}
```

### 2）删除数据

```
func TestDeleteData(t *testing.T) {
    cli, err := ethclient.Dial(NodeUrl)
    if err != nil {
        log.Logger.Error(err)
    }
    instance, err := storage.NewStorage(common.HexToAddress(Address), cli)
    if err != nil {
        log.Logger.Error(err)
    }
    auth, err := eth.GenAuth(cli, PrivateKey)
    if err != nil {
        log.Logger.Error(err)
    }
    tx, err := instance.DeleteData(auth, key1)
    if err != nil {
        log.Logger.Error(err)
    }
    fmt.Println("tx Hash:", tx.Hash().String())
}
```

### 3）修改数据

```
func TestUpdateData(t *testing.T) {
    cli, err := ethclient.Dial(NodeUrl)
    if err != nil {
        log.Logger.Error(err)
    }
    instance, err := storage.NewStorage(common.HexToAddress(Address), cli)
    if err != nil {
        log.Logger.Error(err)
    }
	value, err := instance.QueryData(nil, key1)
    if err != nil {
        log.Logger.Error(err)
    }
	if value == "" {
		log.Logger.Error(errors.New("key1 has no corresponding value in the contract and will not be an update operation"))
	}
	if value == value2 {
		log.Logger.Error(errors.New("the value2 to be updated is the same as the existing value, no update is required"))
	}
    auth, err := eth.GenAuth(cli, PrivateKey)
    if err != nil {
        log.Logger.Error(err)
    }
    tx, err := instance.StoreData(auth, key1, value2)
    if err != nil {
        log.Logger.Error(err)
    }
    fmt.Println("tx Hash:", tx.Hash().String())
}
```

### 4）查询数据
```
func TestQueryData(t *testing.T) {
    cli, err := ethclient.Dial(NodeUrl)
    if err != nil {
        log.Logger.Error(err)
    }
    instance, err := storage.NewStorage(common.HexToAddress(Address), cli)
    if err != nil {
        log.Logger.Error(err)
    }
    value, err := instance.QueryData(nil, key1)
    if err != nil {
        log.Logger.Error(err)
    }
    fmt.Println("value:", value)
}
```

### 5）捕获解析StoreData事件示例
```
func TestCaptureStoreDataEvent(t *testing.T) {
	var storageStoreData = make(chan *storage.StorageStoreData, 100)
	wsCli, err := ethclient.Dial(NodeWsUrl)
    if err != nil {
        log.Logger.Error(err)
    }
	wsInstance, err := storage.NewStorage(common.HexToAddress(Address),wsCli)
	if err != nil {
		log.Logger.Error(err)
	}
	go func() {
	hear:
		transfer, err := wsInstance.StorageFilterer.WatchStoreData(nil, storageStoreData, nil)
		if err != nil {
			log.Logger.Error(err)
			goto hear
		}
		defer transfer.Unsubscribe()
		select {
		case errs := <-transfer.Err():
			log.Logger.Warn(err)
			goto hear
		}
	}()
	for {
		select {
		case event := <-storageStoreData:
			fmt.Println("StoreData key:", event.Key)
			fmt.Println("StoreData value:", event.Value)
		}
	}
}
```

### 6）捕获解析DeleteData事件示例
```
func TestCaptureDeleteDataEvent(t *testing.T) {
	var storageDeleteData = make(chan *storage.StorageDeleteData, 100)
	wsCli, err := ethclient.Dial(NodeWsUrl)
    if err != nil {
        log.Logger.Error(err)
    }
	wsInstance, err := storage.NewStorage(common.HexToAddress(Address),wsCli)
	if err != nil {
		log.Logger.Error(err)
	}
	go func() {
	hear:
		transfer, err := wsInstance.StorageFilterer.WatchDeleteData(nil, storageDeleteData, nil)
		if err != nil {
			log.Logger.Error(err)
			goto hear
		}
		defer transfer.Unsubscribe()
		select {
		case errs := <-transfer.Err():
			log.Logger.Warn(err)
			goto hear
		}
	}()
	for {
		select {
		case event := <-storageDeleteData:
			fmt.Println("DeleteData key:", event.Key)
		}
	}
}
```
