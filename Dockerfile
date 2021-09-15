FROM golang as BUILDER
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 go build -o /goNet cmd/goNet/main.go
RUN CGO_ENABLED=0 go build -o /cdl_tool cmd/cdl_tool/main.go

FROM debian:9
RUN apt update && apt install -y wget
RUN wget https://storage.googleapis.com/sunder-networks/public/stratum-bmv2_0.0.1_amd64.deb -O /tmp/stratum-bmv2_0.0.1_amd64.deb
RUN wget https://storage.googleapis.com/sunder-networks/public/stratum-tools_0.0.1_amd64.deb -O /tmp/stratum-tools_0.0.1_amd64.deb

RUN apt update && \
	apt install -y /tmp/stratum-bmv2_0.0.1_amd64.deb /tmp/stratum-tools_0.0.1_amd64.deb && \
	rm /tmp/stratum-bmv2_0.0.1_amd64.deb /tmp/stratum-tools_0.0.1_amd64.deb

COPY --from=BUILDER /goNet /usr/local/bin/goNet
COPY --from=BUILDER /cdl_tool /usr/local/bin/cdl_tool


# p4 artifacts from https://github.com/stratum/stratum/blob/main/stratum/pipelines/main/main.p4
RUN mkdir -p /p4/out
RUN wget https://storage.googleapis.com/sunder-networks/public/stratum-pipelines_main_bmv2.json -O /p4/out/bmv2.json
RUN wget https://storage.googleapis.com/sunder-networks/public/stratum-pipelines_main_p4info.pb.txt -O /p4/out/p4info.pb.txt