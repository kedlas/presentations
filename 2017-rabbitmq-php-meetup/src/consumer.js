'use strict';

const amqp = require('amqplib/callback_api');
const conf = require('./../config');

const Message = require('./message');
const Resequencer = require('./resequencer');
const SequencedMessage = require('./sequenced_message');
const Sorter = require('./sorter');

const consumerType = process.argv.slice(2).join(' ');
const ex = conf.rabbit.exchange;
let queue;
let processMessage;

const printResult = (message) => {
  console.info(`Processed message "${message.getContent()}".`);
}

if (consumerType === 'simple') {
  queue = conf.rabbit.queues.simple;
  processMessage = (receivedMsg) => {
    const msg = new Message(receivedMsg.properties.headers, receivedMsg.content.toString());
    const duration = Number.parseInt(receivedMsg.properties.headers.duration, 10);
    setTimeout(() => {
      printResult(msg);
    }, duration);
  };

} else if (consumerType === 'sorter') {
  queue = conf.rabbit.queues.sorter;
  const sorter = new Sorter();
  processMessage = (receivedMsg) => {
    const msg = new Message(receivedMsg.properties.headers, receivedMsg.content.toString());
    sorter.init(msg.getId());
    const duration = Number.parseInt(receivedMsg.properties.headers.duration, 10);
    setTimeout(() => {
      const sorted = sorter.getMessages(msg.getId(), msg);
      sorted.forEach((item) => {
        printResult(item);
      });
    }, duration);
  };

} else if (consumerType === 'resequencer') {
  queue = conf.rabbit.queues.resequencer;
  const seq = new Resequencer();
  processMessage = (receivedMsg) => {
    const msg = new SequencedMessage(receivedMsg.properties.headers, receivedMsg.content.toString());
    const duration = Number.parseInt(receivedMsg.properties.headers.duration, 10);
    setTimeout(() => {
      const sorted = seq.getSequence(msg);
      sorted.forEach((item) => {
        printResult(item);
      });
    }, duration);
  };

} else {
  console.error('Invalid consumer type.');
  process.exit(0);
}


  
amqp.connect(conf.rabbit.url, (connErr, conn) => {
  if(connErr) {
    console.error(connErr);
    process.exit(0);
  }

  conn.createChannel((chanErr, ch) => {
    if(chanErr) {
      console.error(chanErr);
      process.exit(0);
    }

    ch.prefetch(conf.message_count);
    ch.assertExchange(ex.name, ex.type, {}, (exErr) => {
      if (exErr) {
        console.error(exErr);
        process.exit(0);
      }

      ch.assertQueue(queue.name, {}, function(err, q) {
        ch.bindQueue(queue.name, ex.name, '');
        ch.consume(queue.name, processMessage, {noAck: true});
        console.info(`[*] Waiting for messages in "${queue.name}"`);
      });
    });
  });
});
