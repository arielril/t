use clap::{Parser, ValueEnum};

#[derive(ValueEnum, Copy, Clone, Debug, PartialEq, Eq)]
enum FormatEncode {
    Utf8,
    Hex,
    Binary,
}

impl std::fmt::Display for FormatEncode {
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
            short = 'i',
            long = "in-format",
            default_value_t = FormatEncode::Utf8,
            value_enum
    )]
    in_format: FormatEncode,
    #[arg(
            short = 'f',
            long = "out-format",
            default_value_t = FormatEncode::Utf8,
            num_args= 0..=1,
            value_enum,
    )]
    out_format: FormatEncode,
    #[arg(required = true)]
    input: Vec<String>,
}

impl Cli {
    pub fn exec(&self) -> Result<(), anyhow::Error> {
        let input = self.input[0..].join(" ");

        println!("input: `{}`", input.chars().take(50).collect::<String>());
        println!("format: {} -> {}", self.in_format, self.out_format);

        if self.in_format == self.out_format {
            println!("{}", input);
            return Ok(());
        }

        let formatter: fn(String, FormatEncode) -> Result<String, anyhow::Error>;

        match self.in_format {
            FormatEncode::Utf8 => formatter = from_utf8,
            FormatEncode::Hex => formatter = from_hex,
            FormatEncode::Binary => formatter = from_bin,
        }

        let formatted = formatter(input.clone(), self.out_format)?;

        println!("{}", formatted);
        Ok(())
    }
}

fn from_utf8(input: String, out_format: FormatEncode) -> Result<String, anyhow::Error> {
    match out_format {
        FormatEncode::Utf8 => Ok(input),
        FormatEncode::Hex => Ok(hex::encode(input)),
        FormatEncode::Binary => Ok(input
            .as_bytes()
            .iter()
            .map(|c| format!("{:08b}", c))
            .collect::<String>()),
    }
}

fn from_hex(input: String, out_format: FormatEncode) -> Result<String, anyhow::Error> {
    match out_format {
        FormatEncode::Hex => Ok(input),
        FormatEncode::Utf8 => {
            let decoded = hex::decode(input)?;
            Ok(String::from_utf8(decoded)?)
        }
        FormatEncode::Binary => from_bin(from_hex(input, FormatEncode::Utf8)?, FormatEncode::Utf8),
    }
}

fn from_bin(input: String, out_format: FormatEncode) -> Result<String, anyhow::Error> {
    match out_format {
        FormatEncode::Binary => Ok(input),
        FormatEncode::Utf8 => {
            let bytes: Vec<u8> = input
                .as_bytes()
                .chunks(8)
                .map(|c| {
                    let bit_str = std::str::from_utf8(c).unwrap();
                    Ok::<u8, anyhow::Error>(u8::from_str_radix(bit_str, 2).unwrap())
                })
                .collect::<Result<Vec<u8>, anyhow::Error>>()?;
            Ok(String::from_utf8(bytes)?)
        }
        _ => Err(anyhow::anyhow!("invalid input format")),
    }
}
