**Prerequire**
Understand:
1. UML - https://evergreens.com.ua/ua/articles/uml-diagrams.html
2. Design Patter:
- https://refactoring.guru/ru/design-patterns
- https://golangbyexample.com/all-design-patterns-golang
3. Microservices Pattern - https://microservices.io


**Concept**
The bot that prepare user to read a book that user should upload or find in database of bot.

The user has:
- store knowing words
- book that can read
- study process
- study log

The bot has:
- database which contain:
1. users (describe all needed data about and for user)
2. books (describe data about book like name, popularity etc.)
3. study-process (describe user's study process)
4. study-log (describe deep info about exist study-process)
5. config? (describe info about project)
6. session-history (describe session of client some cache/tmp-data)
7. users-story (describe where a user stopped where from the user came. This is needed for deep understanding where bot should send the user)

<hr/>

**Architecture**
Config - https://github.com/knadh/koanf

DB provider - SQLite or PostgreSQL (in process of deciding)

How the event from Telegram will handle:
1. For response on this question we should ask ourselves what is problem we want to solve
2. The problem is every message should go to Gateway which must choose right controller for this event\message
3. So we should parse message in right way. 
3.1. If the message is command, It will handle as command
3.2. If the message is text, It will parse by reserved text message
3.3. If the message is replay, It will handle as text from replay
4. So the algorithm is next first 

<hr/>

**What framework can handle with this task?**
- it will be a go-routing approach named Regex Routing

<hr/>


**Word type**
should have:
strength - is count of syllables in word 
frequency - in the book\article etc.
language - it needs to use translation in study

**Book type**
should has:
words - it's mean that all words in use

**Translate pkg**
- https://github.com/st3v/translator

**Syllables pkg**
- https://github.com/advancedlogic/go-freeling
- https://github.com/ernestas-poskus/syllables

**Interesting resources**

Libs fro telegram bot:
- https://github.com/alexwbaule/telegramopenwrt
- https://github.com/olebedev/go-tgbot

Approaches to HTTP routing:
- https://benhoyt.com/writings/go-routing/

Keyboard repository:
- https://github.com/gotgbot/keyboard
- **Pagination**
- https://github.com/ivolkoff/golang-telegram-bot-pagination

**Algorithm for game**
1. get a word
2. find synonyms
3. translate every sysnonym and a word
4. show translation of words

**Hard words contain from 3 and more syllables**
- frequency
- Exception
- disagree

**DB migration**
- https://github.com/golang-migrate/migrate

**ORM**
- https://github.com/ent/ent
- https://gorm.io
- https://github.com/gobuffalo/pop
- https://github.com/go-pg/pg

**Language framework**
- https://github.com/pemistahl/lingua-go