const should = require('chai').should();

const Message = require('./../src/sequenced_message');
const Resequencer = require('./../src/resequencer');

/**
 * RUN these tests separately via: $ node_modules/mocha/bin/_mocha test/resequencer.spec.js
 */
describe('Resequencer', () => {
  it('should return current message when they come in correct sequence', () => {
    const seq = new Resequencer();

    let msg = new Message({ job_id: 'a', sequence_id: 1 }, 'data-a1');
    (seq.getSequence(msg))[0].should.deep.equal(msg);

    msg = new Message({ job_id: 'a', sequence_id: 2 }, 'data-a2');
    (seq.getSequence(msg))[0].should.deep.equal(msg);

    msg = new Message({ job_id: 'b', sequence_id: 1 }, 'data-b1');
    (seq.getSequence(msg))[0].should.deep.equal(msg);

    msg = new Message({ job_id: 'a', sequence_id: 3 }, 'data-a3');
    (seq.getSequence(msg))[0].should.deep.equal(msg);

    msg = new Message({ job_id: 'b', sequence_id: 2 }, 'data-b2');
    (seq.getSequence(msg))[0].should.deep.equal(msg);

    // When comes message which was already processed, return empty array
    msg = new Message({ job_id: 'b', sequence_id: 2 }, 'data-b2');
    seq.getSequence(msg).should.deep.equal([]);
  });
  it('should return ordered messages by their sequenceIds - single buffer', () => {
    const seq = new Resequencer();

    const msg3 = new Message({ job_id: 'a', sequence_id: 3 }, 'data-a3');
    seq.getSequence(msg3).should.deep.equal([]);

    const msg2 = new Message({ job_id: 'a', sequence_id: 2 }, 'data-a2');
    seq.getSequence(msg2).should.deep.equal([]);

    const msg1 = new Message({ job_id: 'a', sequence_id: 1 }, 'data-a1');
    let messages = seq.getSequence(msg1);
    messages.should.have.length(3);
    messages[0].should.deep.equal(msg1);
    messages[1].should.deep.equal(msg2);
    messages[2].should.deep.equal(msg3);

    const msg4 = new Message({ job_id: 'a', sequence_id: 4 }, 'data-a4');
    (seq.getSequence(msg4))[0].should.deep.equal(msg4);

    const msg6 = new Message({ job_id: 'a', sequence_id: 6 }, 'data-a6');
    seq.getSequence(msg6).should.deep.equal([]);

    const msg5 = new Message({ job_id: 'a', sequence_id: 5 }, 'data-a5');
    messages = seq.getSequence(msg5);
    messages.should.have.length(2);
    messages[0].should.deep.equal(msg5);
    messages[1].should.deep.equal(msg6);
  });
  it('should return ordered messages by their sequenceIds - multiple buffer', () => {
    const seq = new Resequencer();

    const msgA3 = new Message({ job_id: 'a', sequence_id: 3 }, 'data-a3');
    seq.getSequence(msgA3).should.deep.equal([]);

    const msgA2 = new Message({ job_id: 'a', sequence_id: 2 }, 'data-a2');
    seq.getSequence(msgA2).should.deep.equal([]);

    const msgB2 = new Message({ job_id: 'b', sequence_id: 2 }, 'data-b2');
    seq.getSequence(msgB2).should.deep.equal([]);

    const msgA1 = new Message({ job_id: 'a', sequence_id: 1 }, 'data-a1');
    let messages = seq.getSequence(msgA1);
    messages.should.have.length(3);
    messages[0].should.deep.equal(msgA1);
    messages[1].should.deep.equal(msgA2);
    messages[2].should.deep.equal(msgA3);

    const msgA4 = new Message({ job_id: 'a', sequence_id: 4 }, 'data-a4');
    (seq.getSequence(msgA4))[0].should.deep.equal(msgA4);

    const msgB1 = new Message({ job_id: 'b', sequence_id: 1 }, 'data-b1');
    messages = seq.getSequence(msgB1);
    messages.should.have.length(2);
    messages[0].should.deep.equal(msgB1);
    messages[1].should.deep.equal(msgB2);
  });
  it('should delete buffer after ttl expires', () => {
    // TODO
  });
});
