@0xef97cf4069588836;

interface Controller {
    getPorts      @0 ()               -> (ports :List(Port));
    listStreams   @1 ()               -> (ids :List(UInt16));
    fetchStream   @2 (id :UInt16)     -> (stream :Stream);
    saveStream    @3 (stream :Stream) -> (id :UInt16);
    deleteStream  @4 (id :UInt16)     -> ();
}

interface Port {
    getConfig    @0 ()                                   -> (config :Config);
    setConfig    @1 (config :Config)                     -> ();
    startSend    @2 (ids :List(UInt16))                  -> ();
    waitSend     @3 (timeout :UInt32)                    -> (done :Bool, error :Text);
    # stopSend    @6  () -> ();
    startCapture @4 (file:Text, packetCount :UInt32)     -> ();
    waitCapture  @5 (timeout :UInt32)                    -> (done :Bool, received :UInt32, dropped :UInt32);
    stopCapture  @6 ()                                   -> ();
    # getStats     @5  () -> ();
    # clearStats   @6  () -> ();
    # saveCapture  @7  () -> ();
    struct Config {
        name        @0 :Text;
    }
}

struct Stream {
    id              @0 :UInt16 = 0;
    count           @1 :UInt32 = 1;
    packetsPerSec   @2 :UInt32 = 1;
    layers          @3 :List(Protocol);
}

struct Field {
    value     @0 :Data;
    mode      @1 :UInt8 = 0;
    step      @2 :Data;
    mask      @3 :Data;
    # it's currently not realistic to have a field that varies over more than
    # 2**16 values. Since the number of packets to generate is equal to the
    # Least Common Multiple of all the fields of a layer, we can already reach
    # huge values with a uint16 integer.
    count     @4 :UInt16 = 1;
}

struct Protocol {
    union {
        ethernet2 :group {
            source       @0  :Field = (mode = 0, mask = 0x"FF FF FF FF FF FF", value = 0x"00 00 00 00 00 00");
            destination  @1  :Field = (mode = 0, mask = 0x"FF FF FF FF FF FF", value = 0x"FF FF FF FF FF FF");
            ethernetType @2  :Field = (mode = 0, mask = 0x"FF FF",             value = 0x"08 00")            ;
        }
        ipv4 :group {
            version      @3  :Field = (mode = 0, mask = 0x"FF",                value = 0x"04")               ;
            ihl          @4  :Field = (mode = 0, mask = 0x"FF",                value = 0x"05")               ;
            tos          @5  :Field = (mode = 0, mask = 0x"FF",                value = 0x"00")               ;
            length       @6  :Field = (mode = 4, mask = 0x"FF FF",             value = 0x"00 00")            ;
            id           @7  :Field = (mode = 0, mask = 0x"FF FF",             value = 0x"00 00")            ;
            flags        @8  :Field = (mode = 0, mask = 0x"FF",                value = 0x"00")               ;
            fragOffset   @9  :Field = (mode = 0, mask = 0x"FF FF",             value = 0x"00 00")            ;
            ttl          @10 :Field = (mode = 0, mask = 0x"FF",                value = 0x"FF")               ;
            protocol     @11 :Field = (mode = 4, mask = 0x"FF",                value = 0x"00")               ;
            checksum     @12 :Field = (mode = 4, mask = 0x"FF FF",             value = 0x"00 00")            ;
            source       @13 :Field = (mode = 0, mask = 0x"FF FF FF FF",       value = 0x"00 00 00 00")      ;
            destination  @14 :Field = (mode = 0, mask = 0x"FF FF FF FF",       value = 0x"FF FF FF FF")      ;
            options      @15 :Field = (mode = 4, mask = 0x"00",                value = 0x"00")               ;
            padding      @16 :Field = (mode = 4, mask = 0x"00",                value = 0x"00")               ;
        }
    }
}
