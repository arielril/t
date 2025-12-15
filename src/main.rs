use simplelog::*;

use crate::commands::hktb::Cli;

mod commands;

fn main() -> Result<(), anyhow::Error> {
    set_logger();

    let tb = Cli::parse();

    tb.exec()?;

    Ok(())
}

fn set_logger() {
    let config = ConfigBuilder::new()
        .add_filter_allow(env!("CARGO_BIN_NAME").to_string())
        .build();
    CombinedLogger::init(vec![TermLogger::new(
        LevelFilter::Debug,
        config,
        TerminalMode::Mixed,
        ColorChoice::Auto,
    )])
    .unwrap();
}
