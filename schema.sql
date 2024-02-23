CREATE TABLE News(
    GUID        varchar(250) NOT NULL,
	Title       varchar(200) NOT NULL,
	Link        varchar(400) NOT NULL,
	PdaLink     varchar(400) NOT NULL,
	Description varchar(1000) NOT NULL,
	PubDate     varchar(100) NOT NULL,
	Category    varchar(100) NOT NULL,
	Author      varchar(100) NOT NULL,
    Created_At   TIMESTAMP
);