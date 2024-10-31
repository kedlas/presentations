
/**
 * Resequencer is responsible for outputting messages ordered by their sequence_id
 */
class Resequencer {

  /**
   * Resequencer constructor
   */
  constructor() {
    this.buffer = {};
  }

  /**
   * Pass actual processed message.
   * This method will return the array of messages that continue
   * @param msg
   * @return {*}
   */
  getSequence(msg) {
    const buf = this._getBuffer(msg.getJobId());

    if (msg.getSequenceId() < buf.waitingFor) {
      // This message was already outputted
      return [];
    }

    // Add message to buffer
    buf.messages[msg.getSequenceId()] = msg;

    // 
    if (msg.getSequenceId() > buf.waitingFor) {
      return [];
    }

    // The message for that resequencer was waiting is now being processed,
    // return the sequence of all messages 
    return this._getMessages(buf, msg);
  }

  /**
   * Retuns existing or creates new buffer
   *
   * @param jobId {string}
   * @private
   */
  _getBuffer(jobId) {
    if (!this.buffer[jobId]) {
      this._createBuffer(jobId);
    }

    return this.buffer[jobId];
  }

  /**
   * 
   * @param {string} jobId 
   * @private
   */
  _createBuffer(jobId) {
    this.buffer[jobId] = {
      messages: {},
      waitingFor: 1,
    };
  }

  /**
   * 
   * @param {object} buffer
   * @param {object} message
   * @private
   */
  _getMessages(buffer, msg) {
    const out = [];
    while (true) {
      const desired = buffer.messages[buffer.waitingFor];
      if (!desired) {
        break;
      }

      out.push(desired);
      delete buffer.messages[buffer.waitingFor];
      buffer.waitingFor += 1;
    }

    return out;
  }

}

module.exports = Resequencer;
