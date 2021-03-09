import requests
from bs4 import BeautifulSoup

DOC_URL: str = "https://docs.docker.com/engine/reference/commandline/run/"

if __name__ == '__main__':
    res = requests.get(DOC_URL).text
    bs = BeautifulSoup(res, features="lxml")

    # find table
    table = bs.find("table")
    # print("table:", table)

    # loop through table rows
    for j, a in enumerate(table.find_all("tr")):
        # skip header
        if j == 0:
            continue

        data = a.find_all("td")

        name: str = data[0].get_text()
        default: str = data[1].get_text()
        desc: str = data[2].get_text()
        aliases: list = []

        # check for aliases
        if "," in name:
            for i, val in enumerate(name.split(",")):
                val = val.strip()
                while val.startswith("-"):
                    val = val[1:]
                if i == 0:
                    name = val
                    continue
                aliases.append(val)

        # update values
        while name.startswith("-"):
            name = name[1:]
        desc = desc.replace("\n", r"\\n")

        name = name.replace('"', r'\"')
        default = default.replace('"', r'\"')
        desc = desc.replace('"', r'\"')

        # generate command
        command: str = "&cli.StringFlag{\n"
        # add name
        command += f'  Name: "{name}",\n'
        # add default
        if len(default) > 0:
            command += f'  Value: "{default}",\n'
        # add desc
        if len(desc) > 0:
            command += f'  Usage: "{desc}",\n'
        # add aliases
        if len(aliases) > 0:
            command += "  Aliases: []string{\n"
            for alias in aliases:
                command += f'    "{alias}",\n'
            command += "  },\n"
        command += "},"

        # print command
        print(command)