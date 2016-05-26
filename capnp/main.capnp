# using Go = import "../../../../zombiezen.com/go/capnproto2/go.capnp";
using Go = import "go.capnp";
$Go.package("capnp");
$Go.import("github.com/little-dude/tgen/capnp");
@0xef97cf4069588836;

interface Controller {
    getPorts @0 () -> (ports :List(Port));
}

interface Port {
    getConfig  @0 ()               -> (config :Config);
    setConfig  @1 (config :Config) -> ();
    getStreams @2 ()               -> (streams :List(Stream));
    newStream  @3 ()               -> (stream :Stream);
    delStream  @4 (name :Text)     -> ();
    # startSend    @1  () -> ();
    # stopSend     @2  () -> ();
    # startCapture @3  () -> ();
    # stopCapture  @4  () -> ();
    # getStats     @5  () -> ();
    # clearStats   @6  () -> ();
    # saveCapture  @7  () -> ();
    struct Config {
        name        @0 :Text;
    }
}

interface Stream {
    getConfig @0 ()                       -> (config :Config);
    setConfig @1 (config :Config)         -> ();
    # getLayers @2 ()                       -> (layers :List(Protocol));
    # setLayers @3 (layers :List(Protocol)) -> ();
    struct Config {
        name            @0 :Text;
        loop            @1 :Bool;
        repeat          @2 :UInt32;
        packetsPerSec   @3 :UInt32;
    }
}

struct Field {
    value     @0 :Data;
    mode      @1 :Text;
    step      @2 :Data;
    mask      @3 :Data;
    count     @4 :UInt64;
}

struct Protocol {
    union {
        ethernet2 :group {
            source       @0  :Field = (mode = "fixed", mask = 0x"FF FF FF FF FF FF", value = 0x"00 00 00 00 00 00");
            destination  @1  :Field = (mode = "fixed", mask = 0x"FF FF FF FF FF FF", value = 0x"FF FF FF FF FF FF");
            ethernetType @2  :Field = (mode = "fixed", mask = 0x"FF FF",             value = 0x"08 00")            ;
        }
        ipv4 :group {
            version      @3  :Field = (mode = "fixed", mask = 0x"FF",                value = 0x"04")               ;
            ihl          @4  :Field = (mode = "fixed", mask = 0x"FF",                value = 0x"05")               ;
            tos          @5  :Field = (mode = "fixed", mask = 0x"FF",                value = 0x"00")               ;
            length       @6  :Field = (mode = "auto" , mask = 0x"FF",                value = 0x"00")               ;
            id           @7  :Field = (mode = "fixed", mask = 0x"FF",                value = 0x"00")               ;
            flags        @8  :Field = (mode = "fixed", mask = 0x"FF",                value = 0x"00")               ;
            fragOffset   @9  :Field = (mode = "fixed", mask = 0x"FF FF",             value = 0x"00 00")            ;
            ttl          @10 :Field = (mode = "fixed", mask = 0x"FF",                value = 0x"FF")               ;
            protocol     @11 :Field = (mode = "auto" , mask = 0x"FF",                value = 0x"00")               ;
            checksum     @12 :Field = (mode = "auto" , mask = 0x"FF",                value = 0x"00")               ;
            srcip        @13 :Field = (mode = "fixed", mask = 0x"FF FF FF FF",       value = 0x"00 00 00 00")      ;
            dstip        @14 :Field = (mode = "fixed", mask = 0x"FF FF FF FF",       value = 0x"FF FF FF FF")      ;
            options      @15 :Field = (mode = "fixed", mask = 0x"FF",                value = 0x"00")               ;
            padding      @16 :Field = (mode = "fixed", mask = 0x"FF",                value = 0x"00")               ;
        }
    }
}
