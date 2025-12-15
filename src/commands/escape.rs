use clap::{Parser, ValueEnum};

#[derive(ValueEnum, Copy, Clone, Debug, PartialEq, Eq)]
enum Escapers {
    Url,
    Html,
}

impl std::fmt::Display for Escapers {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        self.to_possible_value()
            .expect("default serializer shouldn't fail")
            .get_name()
            .fmt(f)
    }
}

#[derive(Parser, Debug)]
#[command()]
pub struct Cli {
    #[arg(
        short='e',
        default_value_t = Escapers::Url,
        value_enum,
    )]
    encoder: Escapers,
    #[arg(required = true)]
    input: Vec<String>,
}

impl Cli {
    pub fn exec(&self) -> Result<(), anyhow::Error> {
        let input = self.input[0..].join(" ");

        match self.encoder {
            Escapers::Url => println!("{}", urlencoding::encode(input.as_str())),
            Escapers::Html => println!("{}", html_escape::encode_text(input.as_str())),
        }

        Ok(())
    }
}
