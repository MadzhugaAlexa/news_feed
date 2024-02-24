DROP TABLE news;

CREATE TABLE News(
	ID serial primary key,
    GUID        varchar(250) NOT NULL,
	Title       varchar(200) NOT NULL,
	Link        varchar(400) NOT NULL,
	PdaLink     varchar(400) NOT NULL,
	Description text NOT NULL,
	PubDate     bigint,
	Category    varchar(100) NOT NULL,
	Author      varchar(100) NOT NULL,
    Created_At   TIMESTAMP
);