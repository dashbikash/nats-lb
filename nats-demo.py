import asyncio
import nats

async def main():
    nc = await nats.connect()
    nc.close()

asyncio.run(main())