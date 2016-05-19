using Go = import "../../../../zombiezen.com/go/capnproto2/go.capnp";
$Go.package("capnp");
$Go.import("github.com/little-dude/tgen/capnp");
@0xef97cf4069588836;

interface Controller {
    getPorts @0 () -> (ports :List(Port));
}

interface Port {
    getName      @0  () -> (name :Text);
    # startSend    @1  () -> ();
    # stopSend     @2  () -> ();
    # startCapture @3  () -> ();
    # stopCapture  @4  () -> ();
    # getStats     @5  () -> ();
    # clearStats   @6  () -> ();
    # saveCapture  @7  () -> ();
    # addStream    @8  (stream :Stream) -> ();
    # delStream    @9  (name :Text) -> ();
    # setStreams   @10  (streams :List(Stream)) -> ();
    # getStreams   @11 () -> (streams :List(Stream));
}

struct Stream {
    name            @0 :Text;
    loop            @1 :Bool;
    repeat          @2 :Int8;
    packetsPerSec   @3 :Int32;
    payload         @4 :Payload;
    layers          @5 :List(Protocol);
}

struct Field {
    enum Mode {
        increment @0;
        decrement @1;
        random    @2;
    }
    mask      @0 :UInt64;
    value     @1 :UInt64;
    step      @2 :UInt64;
    mode      @3 :Mode;
}

struct Payload {
    data        @0 :Data;
    randomize   @1 :Bool;
    length      @2 :UInt32;
}

struct Protocol {
    union {
        ethernet :group {
            source       @0 :Field;
            destination  @1 :Field;
            ethernetType @2 :Field;
            length       @3 :Field;
        }
        ipv4 :group {
            version    @4  :Field;
            ihl        @5  :Field;
            tos        @6  :Field;
            length     @7  :Field;
            id         @8  :Field;
            flags      @9  :Field;
            fragOffset @10 :Field;
            ttl        @11 :Field;
            protocol   @12 :Field;
            checksum   @13 :Field;
            srcip      @14 :Field;
            dstip      @15 :Field;
            options    @16 :Field;
            padding    @17 :Field;
        }
    }
}
