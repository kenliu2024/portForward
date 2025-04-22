import asyncdispatch, asyncnet, strutils

const
  listenPort = 8001
  targetAddr = "yz.testapp.shop"
  targetPort = 8000

proc forward(src, dst: AsyncSocket) {.async.} =
  while true:
    try:
      let data = await src.recv(4096)
      if data.len == 0: break
      await dst.send(data)
    except:
      break
  src.close()
  dst.close()

proc handleClient(client: AsyncSocket) {.async.} =
  let target = newAsyncSocket()
  try:
    await target.connect(targetAddr, Port(targetPort))
    asyncCheck forward(client, target)
    asyncCheck forward(target, client)
  except:
    client.close()
    target.close()

proc main() {.async.} =
  let server = newAsyncSocket()
  server.setSockOpt(OptReuseAddr, true)
  server.bindAddr(Port(listenPort))
  server.listen()
  echo "Listening on port ", listenPort, ", forwarding to ", targetAddr, ":", targetPort

  while true:
    let client = await server.accept()
    asyncCheck handleClient(client)

when isMainModule:
  # 直接运行 main()，并让事件循环持续运行
  waitFor main()