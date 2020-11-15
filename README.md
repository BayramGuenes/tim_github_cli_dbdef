SW-Komponente tim_cli_dborga Zweck/Aufgabe
------------------------------------------
Löschen und Anlegen der Datenbank (Schema timlog, timreceiver,timrepo) .
Eingebunden wird dies als Service Container

Einordnung in das Gesamtkonzept (tim_presse)
--------------------------------------------
Anlage der Tim-Datenbanken.  


Betrieb: Start Komponente
--------------------------------
in configdborga.json sind die maria-db spezifsichen Datenbankeinstellungen (Adresse, Port, Usr) enthalten.
DB-Properties in Volume ./timFileSys/settings/configdborga.json. 



Lokale Ausführung am PC
Start via docker exec -it <ContID> /bin/bash 

<OPT1:> ./cli/db/logger/main confLocation=/cli/config 

<OPT2:> ./cli/db/receiver/main confLocation=/cli/config 

<OPT3:> ./cli/db/repo/main confLocation=/cli/config


Server Ausführung
Start via rancher cli:
<OPT1:> ./cli/db/logger/main confLocation=/cli/config 

<OPT2:> ./cli/db/receiver/main confLocation=/cli/config 

<OPT3:> ./cli/db/repo/main confLocation=/cli/config



UserUI:
Optionen: Eingabe:
'create' zum Anlegen der Datenbank
'drop' zur Löschung der DB
'dropcreate' zum Anlegen der Datenbank mit Löschung der DB
'exit' zum Beenden der Anwendung
-> 

