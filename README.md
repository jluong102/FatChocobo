# Fat Chocobo
Random discord bot

## Usage
Fat Chocobo can be run either locally or through a contianer.
Running the bot in either method will require a config file (See example below).
It is recommend to name the file `config.json` which is the default filename checked.
This can be changed using `-f <filename>` option.
Commands can run by mentioning your bot followed by a command.

### Config Example
```json
{
        "token": "<discord_token>"
}
```

### Running locall
```bash
	make
	./fatchocobo
```

### Running in docker
```bash
	make docker
	docker run -it --rm fatchocobo
```

## Commands
### General 

### Helldivers 2 
* `helldive`: 
	* `planets`: List active campaign planets.
		* `<planet name>`: List information for provided planet.
