## 生成绑定文件

- 生成绑定go文件命令

#### 生成userLocation合约代码

```shell
# Fee
./abigen --abi ./abi/Fee.json --pkg corgi --type Fee --out ./pkg/binding/corgi/fee.go
```

```shell
# Game
./abigen --abi ./abi/Game.json --pkg corgi --type Game --out ./pkg/binding/corgi/game.go
```

```shell
# GameCategory
./abigen --abi ./abi/GameCategory.json --pkg corgi --type GameCategory --out ./pkg/binding/corgi/game_category.go
```

```shell
# Order
./abigen --abi ./abi/Order.json --pkg corgi --type Order --out ./pkg/binding/corgi/order.go
```

```shell
# Token
./abigen --abi ./abi/Token.json --pkg corgi --type Token --out ./pkg/binding/corgi/token.go
```
```shell
# AutoBet
./abigen --abi ./abi/AutoBet.json --pkg corgi --type AutoBet --out ./pkg/binding/corgi/auto_bet.go
```

```shell
# Jackpot
./abigen --abi ./abi/Jackpot.json --pkg corgi --type Jackpot --out ./pkg/binding/corgi/jackpot.go
```



