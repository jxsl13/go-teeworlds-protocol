package snapshot7_test

import (
	"testing"

	"github.com/teeworlds-go/go-teeworlds-protocol/internal/testutils/require"
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/object7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

// snapshot captured with tcpdump
// 0.7 vanilla based hacking on protocol client
// connection to official unique race servers
//
// map hotrun
// server is empty
//

func TestUniqueRaceSnapshot(t *testing.T) {
	t.Parallel()
	// snapshot captured with tcpdump
	// generated by a custom race teeworlds 0.7 server
	// used https://github.com/ChillerDragon/teeworlds/tree/hacking-on-protocol client to connect
	// 0.7 vanilla based client with debug prints
	//
	// verified with https://twnet.zillyhuhn.com/
	//
	// [twnet_parser v0.10.0][huffman=rust-libtw2] udp payload: 10 04 01 3...
	// --- 0.7
	// {
	//   "version": "0.7",
	//   "payload_raw": "16da4d19b6...",
	//   "payload_decompressed": "03261194c30195c301a2e03...",
	//   "header": {
	//     "flags": [
	//       "compression"
	//     ],
	//     "ack": 4,
	//     "token": "3dcc21bb",
	//     "num_chunks": 1
	//   },
	//   "messages": [
	//     {
	//       "message_type": "system",
	//       "message_name": "snap_single",
	//       "system_message": true,
	//       "message_id": 8,
	//       "header": {
	//         "version": "0.7",
	//         "flags": [],
	//         "size": 230,
	//         "seq": -1
	//       },
	//       "tick": 12500,
	//       "delta_tick": 12501,
	//       "crc": 489506,
	//       "data": "0019000413b0cd02b0200..."
	//     }
	//   ]
	// }
	//
	//              No. Time      Source Destination Protocol Length  Info
	// 1st snap ->  19  1.365232  8314   53883       UDP      246     8314 → 53883 Len=204
	//
	dump := []byte{
		0x10, 0x04, 0x01, 0x3d, 0xcc, 0x21,
		0xbb, 0x16, 0xda, 0x4d, 0x19, 0xb6, 0x8c, 0x8a, 0xd9, 0x32, 0xfa, 0xcc, 0xb1, 0x74, 0xa6, 0xc2,
		0x22, 0xcf, 0x9e, 0x2a, 0x58, 0xbe, 0x15, 0xb0, 0xac, 0x7b, 0xf0, 0xb0, 0x8a, 0x2f, 0x60, 0x59,
		0xf7, 0xdc, 0xc1, 0x42, 0x99, 0x80, 0x65, 0xdd, 0x53, 0xc3, 0x72, 0x92, 0x80, 0x65, 0xdd, 0x73,
		0x60, 0x6d, 0x41, 0xc0, 0xb2, 0xee, 0xd1, 0xc3, 0xfa, 0xaa, 0x80, 0x65, 0xdd, 0x0b, 0xcb, 0x63,
		0x02, 0x96, 0x75, 0xcf, 0x06, 0xd6, 0x87, 0x04, 0x2c, 0xeb, 0x9e, 0x80, 0xe5, 0x31, 0x84, 0xf5,
		0x0d, 0x3d, 0xb9, 0x49, 0x1e, 0x43, 0x58, 0xdf, 0xd0, 0xb3, 0xc3, 0xfa, 0x10, 0xc2, 0xfa, 0x86,
		0x9e, 0x85, 0x49, 0x1f, 0x42, 0x58, 0xdf, 0xd0, 0x73, 0x83, 0xd5, 0x35, 0x84, 0xf5, 0x0d, 0x3d,
		0xbd, 0x49, 0x5d, 0x43, 0x58, 0xdf, 0xd0, 0xf3, 0x60, 0x39, 0x19, 0x61, 0x7d, 0x43, 0x8f, 0x98,
		0xe4, 0x64, 0x84, 0xf5, 0x0d, 0x3d, 0x26, 0xac, 0x2d, 0x23, 0xac, 0x6f, 0xe8, 0xf9, 0x9b, 0xb4,
		0x65, 0x84, 0xf5, 0x0d, 0x3d, 0x84, 0xd5, 0x75, 0x84, 0xf5, 0x4d, 0xe8, 0x1b, 0xb6, 0x8c, 0xb0,
		0xaa, 0xa8, 0xd0, 0x08, 0xc4, 0x95, 0x86, 0x55, 0x05, 0x4b, 0x63, 0xfd, 0x78, 0x9b, 0x87, 0xb4,
		0x80, 0x49, 0x83, 0x08, 0x0b, 0x3b, 0xa9, 0x95, 0x56, 0x73, 0xc3, 0xaa, 0x82, 0xa5, 0xe1, 0x3f,
		0x4b, 0x28, 0x92, 0x57, 0xdc, 0x00,
	}

	// slog.SetLogLoggerLevel(slog.LevelDebug)

	packet := protocol7.Packet{}
	err := packet.Unpack(dump)
	require.NoError(t, err)

	// TODO: not working yet
	// conn := protocol7.Session{}
	// conn.Ack = packet.Header.Ack
	// repack := packet.Pack(&conn)
	// require.Equal(t, dump, repack)

	// content
	require.Equal(t, 1, len(packet.Messages))
	require.Equal(t, network7.MsgSysSnapSingle, packet.Messages[0].MsgId())
	msg, ok := packet.Messages[0].(*messages7.SnapSingle)
	require.Equal(t, true, ok)
	require.Equal(t, 12500, msg.GameTick)
	require.Equal(t, 12501, msg.DeltaTick)
	require.Equal(t, 489506, msg.Crc)

	// verified with hacking on protocol print
	require.Equal(t, 25, msg.Snapshot.NumItemDeltas)
	require.Equal(t, 0, msg.Snapshot.NumRemovedItems)
	require.Equal(t, 25, len(msg.Snapshot.Items))

	// TODO: we are missing 2 in the crc
	//       not sure whats going on here
	require.Equal(t, 489504, msg.Snapshot.Crc)

	// verified with hacking on protocol
	item := msg.Snapshot.Items[0]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok := item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 19, pickup.Id())
	require.Equal(t, 21360, pickup.X)
	require.Equal(t, 2096, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[1]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 20, pickup.Id())
	require.Equal(t, 21296, pickup.X)
	require.Equal(t, 2096, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[2]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 15, pickup.Id())
	require.Equal(t, 21232, pickup.X)
	require.Equal(t, 2096, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[3]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 16, pickup.Id())
	require.Equal(t, 21168, pickup.X)
	require.Equal(t, 2096, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[4]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 13, pickup.Id())
	require.Equal(t, 21104, pickup.X)
	require.Equal(t, 2096, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[5]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 14, pickup.Id())
	require.Equal(t, 21040, pickup.X)
	require.Equal(t, 2096, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[6]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 0, pickup.Id())
	require.Equal(t, 20976, pickup.X)
	require.Equal(t, 2096, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[7]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 6, pickup.Id())
	require.Equal(t, 20912, pickup.X)
	require.Equal(t, 2096, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[8]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 2, pickup.Id())
	require.Equal(t, 12784, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[9]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 12, pickup.Id())
	require.Equal(t, 12752, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[10]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 5, pickup.Id())
	require.Equal(t, 12720, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[11]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 3, pickup.Id())
	require.Equal(t, 12688, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[12]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 7, pickup.Id())
	require.Equal(t, 12656, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[13]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 10, pickup.Id())
	require.Equal(t, 12624, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[14]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 4, pickup.Id())
	require.Equal(t, 12592, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[15]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 8, pickup.Id())
	require.Equal(t, 12560, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[16]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 9, pickup.Id())
	require.Equal(t, 12528, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[17]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 11, pickup.Id())
	require.Equal(t, 12496, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[18]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 1, pickup.Id())
	require.Equal(t, 12336, pickup.X)
	require.Equal(t, 1392, pickup.Y)
	require.Equal(t, network7.PickupGrenade, pickup.Type)

	item = msg.Snapshot.Items[19]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok := item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 0, character.Id())
	require.Equal(t, 12500, character.Tick)
	require.Equal(t, 1264, character.X)
	require.Equal(t, 1457, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, 128, character.VelY)
	require.Equal(t, 0, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 1264, character.HookX)
	require.Equal(t, 1456, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 10, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	item = msg.Snapshot.Items[20]
	require.Equal(t, network7.ObjGameData, item.TypeId())
	gameData, ok := item.(*object7.GameData)
	require.Equal(t, true, ok)
	require.Equal(t, 0, gameData.Id())
	require.Equal(t, 0, gameData.GameStartTick)
	require.Equal(t, 0, gameData.FlagsRaw)
	require.Equal(t, 0, gameData.GameStateEndTick)

	item = msg.Snapshot.Items[21]
	require.Equal(t, network7.ObjGameDataRace, item.TypeId())
	gameDataRace, ok := item.(*object7.GameDataRace)
	require.Equal(t, true, ok)
	require.Equal(t, 0, gameDataRace.Id())
	require.Equal(t, 22995, gameDataRace.BestTime)
	require.Equal(t, 3, gameDataRace.Precision)
	require.Equal(t, 5, gameDataRace.RaceFlags)

	item = msg.Snapshot.Items[22]
	require.Equal(t, network7.ObjGameDataFlag, item.TypeId())
	gameDataFlag, ok := item.(*object7.GameDataFlag)
	require.Equal(t, true, ok)
	require.Equal(t, 0, gameDataFlag.Id())
	require.Equal(t, -3, gameDataFlag.FlagCarrierRed)
	require.Equal(t, -3, gameDataFlag.FlagCarrierBlue)
	require.Equal(t, 0, gameDataFlag.FlagDropTickRed)
	require.Equal(t, 0, gameDataFlag.FlagDropTickBlue)

	item = msg.Snapshot.Items[23]
	require.Equal(t, network7.ObjSpawn, item.TypeId())
	spawn, ok := item.(*object7.Spawn)
	require.Equal(t, true, ok)
	require.Equal(t, 0, spawn.Id())
	require.Equal(t, 1264, spawn.X)
	require.Equal(t, 1456, spawn.Y)

	item = msg.Snapshot.Items[24]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok := item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 0, playerInfo.Id())
	require.Equal(t, 0, playerInfo.PlayerFlags)
	require.Equal(t, 105794, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)
}
