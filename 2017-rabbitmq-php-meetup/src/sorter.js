/**
 * Resequencer is responsible for returning messages on the output exactly in the same order,
 * as they were registered(initialized) on the input.
 */
class Sorter {

  /**
   * // TODO - implement sequence/storage size checker
   *
   * Resequencer class constructor
   */
  constructor() {
    this.sequence = [];
    this.storage = {};
  }

  /**
   * Method to be called when received new amqp message.
   * Initializes new record to sequence.
   * @param id
   */
  init(id) {
    this.sequence.push(id);
  }

  /**
   * Method to be called when message was processed.
   * Returns the array of continuous processed messages from the beginning of the sequence.
   * @param id
   * @param data
   * @return {Array}
   */
  getMessages(id, data) {
    this.storage[id] = data;
    let out = [];

    const first = this.getFirst();
    if (first && id === first.id) {
      out = this.getResolved(out);
    }

    return out;
  }

  /**
   * Recursive function to populate out array of values from beginning of the sequence
   * @param out
   * @return {*}
   */
  getResolved(out) {
    const first = this.getFirst();
    if (first !== null && this.storage[first.id]) {
      out.push(this.storage[first.id]);
      delete this.storage[first.id];
      delete this.sequence[first.key];
      out = this.getResolved(out);
    }

    return out;
  }

  /**
   * Returns key, id and value of first element in sequence
   * @return {*}
   */
  getFirst() {
    for (const i in this.sequence) {
      return { key: i, id: this.sequence[i] };
    }

    return null;
  }

}

module.exports = Sorter;
