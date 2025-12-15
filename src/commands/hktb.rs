use clap::{Parser, Subcommand};
use log::error;

#[derive(Parser)]
#[command(name = "tb", author)]
#[command(about = "HKTB - Hacking Tool Belt", long_about = None)]
#[command(before_help = include_str!("banner.txt"))]
#[command(
    version = concat!("v", env!("CARGO_PKG_VERSION")),
    long_version = concat!(include_str!("banner.txt"), "\n", concat!("v", env!("CARGO_PKG_VERSION")))
)]
pub struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Debug, Subcommand)]
enum Commands {
    Decode,
    Decrypt,
    Escape(super::escape::Cli),
    Encrypt,
    Format(super::format::Cli),
}

impl Cli {
    pub fn exec(&self) -> Result<(), anyhow::Error> {
        match &self.command {
            Commands::Format(cmd) => cmd.exec(),
            Commands::Escape(cmd) => cmd.exec(),
            _ => {
                error!("could not find selected command");

                Ok(())
            }
        }
    }

    pub fn parse() -> Self {
        Parser::parse()
    }
}
