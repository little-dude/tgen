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
    getConfig @0 ()               -> (config :Config);
    setConfig @1 (config :Config) -> ();
    # getLayers @2 () -> (layers :List(Protocol));
    # setLayers @3 (layers :List(Protocol)) -> ();
    # enable    @4 () -> ();
    # disable   @5 () -> ();
    struct Config {
        name            @0 :Text;
        loop            @1 :Bool;
        repeat          @2 :UInt32;
        packetsPerSec   @3 :UInt32;
    }
}


# struct Field {
#     value     @0 :Data;
#     step      @1 :Data;
#     mode      @2 :Text;
#     mask      @3 :Data;
#     offset    @4 :UInt8;
# }
# struct Protocol {
#     union {
#         ethernet :group {
#             source       @0 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF FF FF FF FF FF", value = 0x"00 00 00 00 00 00");
#             destination  @1 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF FF FF FF FF FF", value = 0x"FF FF FF FF FF FF");
#             ethernetType @2 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF FF",             value = 0x"08 00")            ;
#             length       @3 :Field;
#         }
#         ipv4 :group {
#             version    @4  :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF",          value = 0x"04")         ;
#             ihl        @5  :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF",          value = 0x"05")         ;
#             tos        @6  :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF",          value = 0x"00")         ;
#             length     @7  :Field;
#             id         @8  :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF",          value = 0x"00")         ;
#             flags      @9  :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF",          value = 0x"00")         ;
#             fragOffset @10 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF FF",       value = 0x"00 00")      ;
#             ttl        @11 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF",          value = 0x"FF")         ;
#             protocol   @12 :Field;
#             checksum   @13 :Field;
#             srcip      @14 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF FF FF FF", value = 0x"00 00 00 00");
#             dstip      @15 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF FF FF FF", value = 0x"FF FF FF FF");
#             options    @16 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF",          value = 0x"00")         ;
#             padding    @17 :Field = (mode = "fixed", step = 0x"00", offset = 0, mask = 0x"FF",          value = 0x"00")         ;
#         }
#     }
# }
struct Field8 {
    value     @0 :UInt8;
    step      @1 :UInt8;
    mode      @2 :Text;
    count     @3 :UInt8;
}

struct Field16 {
    value     @0 :UInt16;
    step      @1 :UInt16;
    mode      @2 :Text;
    count     @3 :UInt16;
}

struct Field32 {
    value     @0 :UInt32;
    step      @1 :UInt32;
    mode      @2 :Text;
    count     @3 :UInt32;
}

struct Field64 {
    value     @0 :UInt64;
    step      @1 :UInt64;
    mode      @2 :Text;
    count     @3 :UInt64;
}

struct LongField {
    value     @0 :Data;
    step      @1 :Data;
    mode      @2 :Text;
    count     @3 :Data;
    mask      @4 :Data;
}

struct Protocol {
    union {
        ethernet :group {
            source       @0 :LongField;
            destination  @1 :LongField;
            ethernetType @2 :Field16;
            length       @3 :Field16;
        }
        ipv4 :group {
            version    @4  :Field8;
            ihl        @5  :Field8;
            tos        @6  :Field8;
            length     @7  :Field8;
            id         @8  :Field8;
            flags      @9  :Field8;
            fragOffset @10 :Field16;
            ttl        @11 :Field8;
            protocol   @12 :Field8;
            checksum   @13 :Field8;
            srcip      @14 :Field32;
            dstip      @15 :Field32;
            options    @16 :LongField;
            padding    @17 :Field8;
        }
    }
}
