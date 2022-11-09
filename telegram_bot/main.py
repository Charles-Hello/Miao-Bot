import json
import re

import requests
from telethon import events, TelegramClient

requests.packages.urllib3.disable_warnings()
import asyncio
import config

client = TelegramClient("user", config.App_api_id, config.App_api_hash,
                        connection_retries=None).start()


@client.on(events.NewMessage(chats=1716089227))
async def check_id(event):
    sender = await event.get_sender()
    # message = await client.get_messages(sender.id, ids=event.message.id)
    sender = await event.get_sender()
    # print("当前聊天的id：" + str(event.chat_id))  # chat_id是int类型
    # print("发送信息：" + event.raw_text)
    # print(type(event.raw_text))
    # print(event.raw_text)
    # print("这个信息发送者的id：" + str(event.sender_id))
    # print("这个信息的id：" + str(event.id))
    # print("是否频道：" + str(event.is_channel))
    # print("是否群组：" + str(event.is_group))

    last_name = "None"
    first_name = "None"
    phone = "None"
    bot = "None"
    chanel_name = "None"
    if str(event.is_channel) != "True":
        phone = re.findall('phone=(.*)', str(sender))[0].split(", ")[0]

    # if 'title' in str(message):
    #     chanel_name = re.findall('title=(.*)', str(sender))[0].split(", ")[0]

    username = re.findall('username=(.*)', str(sender))[0].split(", ")[0]
    try:
        bot = re.findall('bot=(.*)', str(sender))[0].split(", ")[0]
    except:
        pass
    if "title" not in str(sender):
        first_name = re.findall('first_name=(.*)', str(sender))[0].split(", ")[0]
        last_name = re.findall('last_name=(.*)', str(sender))[0].split(", ")[0]
    print("姓：" + first_name)
    print("名：" + last_name)
    print("username名字：" + username)
    print("电话：" + phone)
    print("是否机器人：" + bot)
    print("频道名字：" + chanel_name)
    print("--" * 20)
    data = {
        'Tg_msg_from_id': str(event.sender_id),
        'Tg_msg': str(event.raw_text),
        'Tg_ifbot': bot,
        'Tg_ifgroup': str(event.is_group),
        'Tg_ifchanel': str(event.is_channel),
        'Tg_first_name': first_name,
        'Tg_least_name': last_name,
        'Tg_username': username,
        'Tg_chanel_name': str(chanel_name),
    }
    json_str = json.dumps(data)
    response = requests.post(url=config.API_URL, data=json_str)
    Gin_data = json.loads(response.text)
    OpSlice_list = Gin_data['Response']['OpSlice']
    print(OpSlice_list)
    if OpSlice_list:
        for i in OpSlice_list:
            if 'delete' in i:
                await event.delete()
                await client.send_message(event.chat_id, '我已经被删除')
            elif 'edit' in i:
                edit_message = re.findall('edit(.*)', i)[0]
                await event.edit(edit_message)
            elif 'reply' in i:
                reply_message = re.findall('reply(.*)', i)[0]
                await event.reply(reply_message)
            elif 'send' in i:
                send_message = re.findall('send(.*)', i)[0]
                print(send_message)
                await client.send_message(event.chat_id, send_message)
            elif 'sleep' in i:
                await asyncio.sleep(1.5)


if __name__ == '__main__':
    try:
        asyncio.get_event_loop().run_forever()
    except (KeyboardInterrupt, SystemExit) as e:
        print(e)
