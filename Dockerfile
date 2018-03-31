FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get gopkg.in/tucnak/telebot.v2  
RUN go get github.com/tidwall/gjson
RUN go get gopkg.in/mgo.v2
RUN	go get gopkg.in/mgo.v2/bson
RUN go get build -o main . 
CMD ["/app/main"]