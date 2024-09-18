use std::sync::mpsc::Receiver;
use std::error::Error;

#[derive(Clone, Debug, Eq, Hash, PartialEq)]
pub struct NetAddr(pub String);

pub struct RPC {
    pub from: NetAddr,
    pub payload: Vec<u8>,
}

pub trait Transport {
    fn consume(&self) -> Receiver<RPC>;
    fn connect(&mut self, transport: &dyn Transport) -> Result<(), Box<dyn Error>>;
    fn send_message(&self, address: NetAddr, payload: Vec<u8>) -> Result<(), Box<dyn Error>>;
    fn address(&self) -> NetAddr;
}

